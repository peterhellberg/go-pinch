// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

import (
	"bytes"
	"compress/flate"
	"errors"
	"fmt"
	"unsafe"
)

// Get a file from URL and filename (string)
func Get(url, fn string) ([]byte, error) {
	entries, err := GetZipDirectory(url)

	if err != nil {
		return nil, err
	}

	entry := entries[fn]

	if len(entry.Filename) > 0 {
		file, err := GetZipFile(url, entry)

		if err != nil {
			return nil, err
		}

		return file, nil
	}

	return nil, errors.New("File not found")
}

// Get a file from URL and ZipEntry
func GetZipFile(url string, entry ZipEntry) ([]byte, error) {
	echoZipEntry(&entry)

	// Using hardcoded 30 as go length include some padding,
	// 16 added because seen extraFieldLength differ between
	// file header and directory entry
	length := 30 + entry.CompressedSize + uint32(len(entry.Filename)) + uint32(entry.ExtraFieldLength) + 16

	o := int64(entry.RelativeOffsetOfLocalFileHeader)

	echo("\nGet Zip File:")
	body, err := fetchPartialData(url, o, o+int64(length-1))

	if err != nil {
		return nil, err
	}

	var file *ZipFileHeader
	file = (*ZipFileHeader)(unsafe.Pointer(&body[0]))

	echo(594828 - o)

	if file.LocalFileHeaderSignature == 0x04034b50 {
		offset := file.StartOffset()

		if file.CompressionMethod == 8 {
			data := body[offset : offset+file.CompressedSize()]

			zipreader := flate.NewReader(bytes.NewReader(data))

			buf := new(bytes.Buffer)
			buf.ReadFrom(zipreader)

			b := buf.Bytes()
			zipreader.Close()

			return b, nil
		} else if file.CompressionMethod == 0 {
			return body[offset : offset+file.OriginalSize()], nil
		}

		err = errors.New("Unimplemented compression method")
	} else {
		err = errors.New("Corrupt file (signature error)")
	}

	return nil, err
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

	echo("\nGet 4k from EOF and look for the End Record:")
	body, err := fetchPartialData(url, of, cl)

	if err != nil {
		return nil, err
	}

	entries := make(map[string]ZipEntry)

	var dir *ZipDirRecord
	var rec *ZipEndRecord

	// Find the End Record
	endOffset := bytes.Index(body, []byte{0x50, 0x4b, 0x05, 0x06})

	if endOffset >= 0 {
		buf := body[endOffset : endOffset+int(unsafe.Sizeof(ZipEndRecord{}))]

		rec = (*ZipEndRecord)(unsafe.Pointer(&buf[0]))

		echo("\nGet the Central Directory Record:")
		body, err := fetchPartialData(url, rec.StartOffset(), rec.EndOffset())

		if err != nil {
			return nil, err
		}

		var l int32 = int32(rec.SizeOfCentralDirectory)
		var i int32 = 0

		// Read the entries
		for l > 46 {
			buf = body[i : i+int32(unsafe.Sizeof(ZipDirRecord{}))]

			dir = (*ZipDirRecord)(unsafe.Pointer(&buf[0]))

			if dir.CentralDirectoryFileHeaderSignature == 0x02014b50 {
				var entry ZipEntry

				// Only collect files (skipping directories)
				if (dir.ExternalFileAttributesH & 0x4000) != 0x4000 {
					fn := string(body[i+46 : i+46+int32(dir.FileNameLength)])

					populateEntry(&entry, dir, fn)

					entries[fn] = entry
				}

				l = l - dir.CombinedLength()
				i = i + dir.CombinedLength()
			} else {
				err = errors.New("Corrupt directory (signature error)")
				break
			}
		}
	}

	return entries, err
}

func populateEntry(entry *ZipEntry, dir *ZipDirRecord, fn string) *ZipEntry {
	entry.Filename = fn
	entry.CompressedSize = dir.CompressedSize
	entry.UncompressedSize = dir.UncompressedSize
	entry.CompressionMethod = dir.CompressionMethod
	entry.ExtraFieldLength = dir.ExtraFieldLength
	entry.RelativeOffsetOfLocalFileHeader = dir.RelativeOffset()

	return entry
}

func echoZipEntry(entry *ZipEntry) {
	echo("\nZipEntry:")
	echo(" Filename                       ", entry.Filename)
	echo(" CompressedSize                 ", entry.CompressedSize)
	echo(" UncompressedSize               ", entry.UncompressedSize)
	echo(" CompressionMethod              ", entry.CompressionMethod)
	echo(" ExtraFieldLength               ", entry.ExtraFieldLength)
	echo(" RelativeOffsetOfLocalFileHeader", entry.RelativeOffsetOfLocalFileHeader)
}

func echo(v ...interface{}) {
	debug := true

	if debug {
		fmt.Println(v...)
	}
}
