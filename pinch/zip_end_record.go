// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

import (
	"fmt"
)

// ZipEndRecord zip archive end record, see http://en.wikipedia.org/wiki/ZIP_(file_format)#File_headers
type ZipEndRecord struct {
	EndOfCentralDirectorySignature            uint32
	NumberOfThisDisk                          uint16
	DiskWhereCentralDirectoryStarts           uint16
	NumberOfCentralDirectoryRecordsOnThisDisk uint16
	TotalNumberOfCentralDirectoryRecords      uint16
	SizeOfCentralDirectory                    uint32
	OffsetOfStartOfCentralDirectory           uint32
	ZipfileCommentLength                      uint16
}

// StartOffset get the offset of the zip archive directory
func (r *ZipEndRecord) StartOffset() int64 {
	return int64(r.OffsetOfStartOfCentralDirectory)
}

// EndOffset get the offset of the end of the zip archive directory
func (r *ZipEndRecord) EndOffset() int64 {
	s := r.SizeOfCentralDirectory + r.OffsetOfStartOfCentralDirectory

	return int64(s)
}

func (r *ZipEndRecord) echo() {
	echo("ZipEndRecord")

	echo(" EndOfCentralDirectorySignature           ", fmt.Sprintf("%U", r.EndOfCentralDirectorySignature))
	echo(" NumberOfThisDisk                         ", r.NumberOfThisDisk)
	echo(" DiskWhereCentralDirectoryStarts          ", r.DiskWhereCentralDirectoryStarts)
	echo(" NumberOfCentralDirectoryRecordsOnThisDisk", r.NumberOfCentralDirectoryRecordsOnThisDisk)
	echo(" TotalNumberOfCentralDirectoryRecords     ", r.TotalNumberOfCentralDirectoryRecords)
	echo(" SizeOfCentralDirectory                   ", r.SizeOfCentralDirectory)
	echo(" OffsetOfStartOfCentralDirectory          ", r.OffsetOfStartOfCentralDirectory)
	echo(" ZipfileCommentLength                     ", r.ZipfileCommentLength)
	echo("")
}
