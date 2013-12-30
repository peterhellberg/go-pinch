// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

import (
	"math"
)

type ZipDirRecord struct {
	CentralDirectoryFileHeaderSignature uint32
	VersionMadeBy                       uint16
	VersionNeededToExtract              uint16
	GeneralPurposeBitFlag               uint16
	CompressionMethod                   uint16
	FileLastModificationTime            uint16
	FileLastModificationDate            uint16
	Crc32                               uint32
	CompressedSize                      uint32
	UncompressedSize                    uint32
	FileNameLength                      uint16
	ExtraFieldLength                    uint16
	FileCommentLength                   uint16
	DiskNumberWhereFileStarts           uint16
	InternalFileAttributes              uint16
	ExternalFileAttributesL             uint16 // split in low+high for struct packing
	ExternalFileAttributesH             uint16
	RelativeOffsetOfLocalFileHeaderL    uint16 // split in low+high for struct packing
	RelativeOffsetOfLocalFileHeaderH    uint16
}

func (d *ZipDirRecord) RelativeOffset() uint32 {
	l := d.RelativeOffsetOfLocalFileHeaderL
	h := d.RelativeOffsetOfLocalFileHeaderH * math.MaxUint16

	return uint32(l + h)
}

func (d *ZipDirRecord) CombinedLength() int32 {
	l := d.FileNameLength + d.ExtraFieldLength + d.FileCommentLength

	return int32(46 + l)
}
