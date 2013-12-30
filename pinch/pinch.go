// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

import (
	"bytes"
	"compress/flate"
	"errors"
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
	// Using hardcoded 30 as go length include some padding,
	// 16 added because seen extraFieldLength differ between
	// file header and directory entry
	length := 30 + entry.compressedSize + uint32(len(entry.Filename)) + uint32(entry.extraFieldLength) + 16

	body, err := fetchPartialData(url, int64(entry.relativeOffsetOfLocalFileHeader), int64(length-1))

	if err != nil {
		return nil, err
	}

	var file *ZipFileHeader
	file = (*ZipFileHeader)(unsafe.Pointer(&body[0]))

	if file.localFileHeaderSignature == 0x04034b50 {
		offset := file.StartOffset()

		if file.compressionMethod == 8 {
			data := body[offset : offset+file.CompressedSize()]

			zipreader := flate.NewReader(bytes.NewReader(data))

			buf := new(bytes.Buffer)
			buf.ReadFrom(zipreader)

			b := buf.Bytes()
			zipreader.Close()

			return b, nil
		} else if file.compressionMethod == 0 {
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

	// Get 4k from end of file and look for the End Record
	body, err := fetchPartialData(url, of, cl)

	if err != nil {
		return nil, err
	}

	entries := make(map[string]ZipEntry)

	// Find the End Record
	endoffset := bytes.Index(body, []byte{0x50, 0x4b, 0x05, 0x06})

	if endoffset >= 0 {
		buf := body[endoffset : endoffset+int(unsafe.Sizeof(ZipEndRecord{}))]

		var rec *ZipEndRecord
		rec = (*ZipEndRecord)(unsafe.Pointer(&buf[0]))

		// Get the Central Directory Record
		body, err := fetchPartialData(url, rec.StartOffset(), rec.EndOffset())

		if err != nil {
			return nil, err
		}

		var l int32 = int32(rec.sizeOfCentralDirectory)
		var i int32 = 0

		// Read the entries
		for l > 46 {
			buf = body[i : i+int32(unsafe.Sizeof(ZipDirRecord{}))]

			var dir *ZipDirRecord
			dir = (*ZipDirRecord)(unsafe.Pointer(&buf[0]))

			if dir.centralDirectoryFileHeaderSignature == 0x02014b50 {
				var entry ZipEntry

				// Only collect files (skipping directories)
				if (dir.externalFileAttributesH & 0x4000) != 0x4000 {
					fn := string(body[i+46 : i+46+int32(dir.fileNameLength)])

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
	entry.compressedSize = dir.compressedSize
	entry.uncompressedSize = dir.uncompressedSize
	entry.compressionMethod = dir.compressionMethod
	entry.extraFieldLength = dir.extraFieldLength
	entry.relativeOffsetOfLocalFileHeader = dir.RelativeOffset()

	return entry
}
