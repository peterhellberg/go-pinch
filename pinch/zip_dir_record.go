// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

import (
	"math"
)

type ZipDirRecord struct {
	centralDirectoryFileHeaderSignature uint32
	versionMadeBy                       uint16
	versionNeededToExtract              uint16
	generalPurposeBitFlag               uint16
	compressionMethod                   uint16
	fileLastModificationTime            uint16
	fileLastModificationDate            uint16
	crc32                               uint32
	compressedSize                      uint32
	uncompressedSize                    uint32
	fileNameLength                      uint16
	extraFieldLength                    uint16
	fileCommentLength                   uint16
	diskNumberWhereFileStarts           uint16
	internalFileAttributes              uint16
	externalFileAttributesL             uint16 // split in low+high for struct packing
	externalFileAttributesH             uint16
	relativeOffsetOfLocalFileHeaderL    uint16 // split in low+high for struct packing
	relativeOffsetOfLocalFileHeaderH    uint16
}

func (d *ZipDirRecord) RelativeOffset() uint32 {
	l := d.relativeOffsetOfLocalFileHeaderL
	h := d.relativeOffsetOfLocalFileHeaderH * math.MaxUint16

	return uint32(l + h)
}

func (d *ZipDirRecord) CombinedLength() int32 {
	l := d.fileNameLength + d.extraFieldLength + d.fileCommentLength

	return int32(46 + l)
}
