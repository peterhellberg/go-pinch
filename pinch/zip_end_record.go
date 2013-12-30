// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

type ZipEndRecord struct {
	endOfCentralDirectorySignature            uint32
	numberOfThisDisk                          uint16
	diskWhereCentralDirectoryStarts           uint16
	numberOfCentralDirectoryRecordsOnThisDisk uint16
	totalNumberOfCentralDirectoryRecords      uint16
	sizeOfCentralDirectory                    uint32
	offsetOfStartOfCentralDirectory           uint32
	zipfileCommentLength                      uint16
}

func (r *ZipEndRecord) StartOffset() int64 {
	return int64(r.offsetOfStartOfCentralDirectory)
}

func (r *ZipEndRecord) EndOffset() int64 {
	s := r.sizeOfCentralDirectory + r.offsetOfStartOfCentralDirectory - 1

	return int64(s)
}
