/*---------------------------------------------------------------------------
 
 Copyright (c) 2013 Edward Patel
 
 Permission is hereby granted, free of charge, to any person obtaining
 a copy of this software and associated documentation files (the
 "Software"), to deal in the Software without restriction, including
 without limitation the rights to use, copy, modify, merge, publish,
 distribute, sublicense, and/or sell copies of the Software, and to
 permit persons to whom the Software is furnished to do so, subject to
 the following conditions:

 The above copyright notice and this permission notice shall be
 included in all copies or substantial portions of the Software.

 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
 LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
 OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 
 ---------------------------------------------------------------------------*/

package pinch

import (
	"bytes"
	"fmt"
	"math"
	"errors"
	"unsafe"
	"net/http"
	"io/ioutil"
	"compress/flate"
)

type zip_end_record struct {
	endOfCentralDirectorySignature				uint32
	numberOfThisDisk							uint16
	diskWhereCentralDirectoryStarts				uint16
	numberOfCentralDirectoryRecordsOnThisDisk	uint16
	totalNumberOfCentralDirectoryRecords		uint16
	sizeOfCentralDirectory						uint32
	offsetOfStartOfCentralDirectory				uint32
	zipfileCommentLength						uint16
}

type zip_dir_record struct {
	centralDirectoryFileHeaderSignature	uint32
	versionMadeBy						uint16
	versionNeededToExtract				uint16
	generalPurposeBitFlag				uint16
	compressionMethod					uint16
	fileLastModificationTime			uint16
	fileLastModificationDate			uint16
	crc32								uint32
	compressedSize						uint32
	uncompressedSize					uint32
	fileNameLength						uint16
	extraFieldLength					uint16
	fileCommentLength					uint16
	diskNumberWhereFileStarts			uint16
	internalFileAttributes				uint16
	externalFileAttributesL				uint16 // split in low+high for struct packing
	externalFileAttributesH				uint16
	relativeOffsetOfLocalFileHeaderL    uint16 // split in low+high for struct packing
	relativeOffsetOfLocalFileHeaderH    uint16
}

type zip_file_header struct {
    localFileHeaderSignature    uint32
    versionNeededToExtract      uint16
    generalPurposeBitFlag       uint16
    compressionMethod           uint16
    fileLastModificationTime    uint16
    fileLastModificationDate    uint16
    crc32L                      uint16
    crc32H                      uint16
    compressedSizeL             uint16
    compressedSizeH             uint16
    uncompressedSizeL           uint16
    uncompressedSizeH           uint16
    fileNameLength              uint16
    extraFieldLength            uint16
}

type ZipEntry struct {
	Filename                         string 
	compressedSize                   uint32
	uncompressedSize                 uint32
	compressionMethod                uint16
    extraFieldLength                 uint16
	relativeOffsetOfLocalFileHeader  uint32
}

func GetZipDirectory(url string) (map[string]ZipEntry, error) {
	var of int64
	cl, err := getContentLength(url)
	if err != nil {
		return nil, err
	}

	if cl >= 4096 {
		of = cl - 4096
	}

	// Get 4k from end of file and look for the End Record
	body, err := fetchPartialData(url, of, cl)
	if err != nil {
		return nil, err
	}

	entries := make(map[string]ZipEntry)

	// Find the End Record
	endoffset := bytes.Index(body, []byte{ 0x50, 0x4b, 0x05, 0x06 })
	if endoffset >= 0 {
		buf := body[endoffset : endoffset+int(unsafe.Sizeof(zip_end_record{}))]
		var rec *zip_end_record
		rec = (*zip_end_record)(unsafe.Pointer(&buf[0]))		

		// Get the Central Directory Record
		body, err := fetchPartialData(url, int64(rec.offsetOfStartOfCentralDirectory), int64(rec.sizeOfCentralDirectory + rec.offsetOfStartOfCentralDirectory - 1))
		if err != nil {
			return nil, err
		}
		
		var l int16 = int16(rec.sizeOfCentralDirectory)
		var i int32 = 0

		// Read the entries
		for l > 46 {

			var entry ZipEntry
			buf = body[i : i+int32(unsafe.Sizeof(zip_dir_record{}))]
			
			var dir *zip_dir_record
			dir = (*zip_dir_record)(unsafe.Pointer(&buf[0]))
			
			if dir.centralDirectoryFileHeaderSignature == 0x02014b50 {

				if (dir.externalFileAttributesH & 0x4000) != 0x4000 { // Only collect files (skipping directories)
					entry.Filename = string(body[i+46 : i+46+int32(dir.fileNameLength)])
					entry.compressedSize = dir.compressedSize
					entry.uncompressedSize = dir.uncompressedSize
					entry.compressionMethod = dir.compressionMethod
					entry.extraFieldLength = dir.extraFieldLength
					entry.relativeOffsetOfLocalFileHeader = uint32(dir.relativeOffsetOfLocalFileHeaderL + dir.relativeOffsetOfLocalFileHeaderH * math.MaxUint16)
					entries[entry.Filename] = entry
				}
				l = l - (46 + int16(dir.fileNameLength + dir.extraFieldLength + dir.fileCommentLength))
				i = i + int32(46 + dir.fileNameLength + dir.extraFieldLength + dir.fileCommentLength)
			
			} else {
			
				err = errors.New("Corrupt directory (signature error)")
				break
			}
		}
	}

	return entries, err
}

func GetZipFile(url string, entry ZipEntry) ([]byte, error) {

	// Using hardcoded 30 as go length include some padding, 16 added because seen extraFieldLength differ between 
	// file header and directory entry
	length := 30 + entry.compressedSize + uint32(len(entry.Filename)) + uint32(entry.extraFieldLength) + 16

	body, err := fetchPartialData(url, int64(entry.relativeOffsetOfLocalFileHeader), int64(length - 1))
	if err != nil {
		return nil, err
	}

	var file *zip_file_header
	file = (*zip_file_header)(unsafe.Pointer(&body[0]))

	if file.localFileHeaderSignature == 0x04034b50 {

		offset := 30 + uint32(file.fileNameLength) + uint32(file.extraFieldLength)

		if file.compressionMethod == 8 {

			zipreader := flate.NewReader(bytes.NewReader(body[offset : offset+uint32(file.compressedSizeL) + (uint32(file.compressedSizeH) << 16)]))
			buf := new(bytes.Buffer)
			buf.ReadFrom(zipreader)
			b := buf.Bytes()
			zipreader.Close()

			return b, nil

		} else if file.compressionMethod == 0 {

			return body[offset : offset+uint32(file.compressedSizeL + file.compressedSizeH * math.MaxUint16)], nil

		}

		err = errors.New("Unimplemented compression method")

	} else {

		err = errors.New("Corrupt file (signature error)")

	}

	return nil, err
}

func getContentLength(url string) (int64, error) {

	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != 200 {
		return 0, errors.New(resp.Status)
	}

	return resp.ContentLength, nil
}

func fetchPartialData(url string, sof int64, eof int64) ([]byte, error) {

	// Create a client to be used to add the "Range:" header
	client := &http.Client {
		// Go net/http does not keep headers in redirects, add specifically
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return errors.New("To many redirects")
			}
			req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", sof, eof))
			return nil
		},
	}

	// Create a request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add the "Range:" header
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", sof, eof))

	// Now run the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Clean up	
	defer resp.Body.Close()

	// Get body
	body, err := ioutil.ReadAll(resp.Body)

	// And return it...
	return body, err
}
