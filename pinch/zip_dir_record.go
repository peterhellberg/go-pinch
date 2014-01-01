// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

import (
	"fmt"
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

func (d *ZipDirRecord) CombinedLength() int32 {
	l := d.FileNameLength + d.ExtraFieldLength + d.FileCommentLength

	return int32(46 + l)
}

func (d *ZipDirRecord) ExternalFileAttributes() uint32 {
	l := uint32(d.ExternalFileAttributesL)
	h := (uint32(d.ExternalFileAttributesH) << 16)

	return l + h
}

func (d *ZipDirRecord) RelativeOffset() uint32 {
	l := uint32(d.RelativeOffsetOfLocalFileHeaderL)
	h := (uint32(d.RelativeOffsetOfLocalFileHeaderH) << 16)

	return l + h
}

func (d *ZipDirRecord) echo() {
	echo("ZipDirRecord")
	echo(" CentralDirectoryFileHeaderSignature", fmt.Sprintf("%U", d.CentralDirectoryFileHeaderSignature))
	echo(" VersionMadeBy                      ", d.VersionMadeBy)
	echo(" VersionNeededToExtract             ", d.VersionNeededToExtract)
	echo(" GeneralPurposeBitFlag              ", d.GeneralPurposeBitFlag)
	echo(" CompressionMethod                  ", d.CompressionMethod)
	echo(" FileLastModificationTime           ", d.FileLastModificationTime)
	echo(" FileLastModificationDate           ", d.FileLastModificationDate)
	echo(" Crc32                              ", d.Crc32)
	echo(" CompressedSize                     ", d.CompressedSize)
	echo(" UncompressedSize                   ", d.UncompressedSize)
	echo(" FileNameLength                     ", d.FileNameLength)
	echo(" ExtraFieldLength                   ", d.ExtraFieldLength)
	echo(" FileCommentLength                  ", d.FileCommentLength)
	echo(" DiskNumberWhereFileStarts          ", d.DiskNumberWhereFileStarts)
	echo(" InternalFileAttributes             ", d.InternalFileAttributes)
	echo(" ExternalFileAttributes()           ", d.ExternalFileAttributes())
	echo(" RelativeOffset()                   ", d.RelativeOffset())
}
