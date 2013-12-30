// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

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

func (r *ZipEndRecord) StartOffset() int64 {
	return int64(r.OffsetOfStartOfCentralDirectory)
}

func (r *ZipEndRecord) EndOffset() int64 {
	s := r.SizeOfCentralDirectory + r.OffsetOfStartOfCentralDirectory - 1

	return int64(s)
}
