// Copyright (c) 2013-2014 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

import (
	"fmt"
)

// ZipDirRecord zip archive directory record, see http://en.wikipedia.org/wiki/ZIP_(file_format)#File_headers
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
	FilenameLength                      uint16
	ExtraFieldLength                    uint16
	FileCommentLength                   uint16
	DiskNumberWhereFileStarts           uint16
	InternalFileAttributes              uint16
	ExternalFileAttributesL             uint16 // split in low+high for struct packing
	ExternalFileAttributesH             uint16
	RelativeOffsetOfLocalFileHeaderL    uint16 // split in low+high for struct packing
	RelativeOffsetOfLocalFileHeaderH    uint16
}

// CombinedLength calculate the combined length of the record + filename + extra fields
func (d *ZipDirRecord) CombinedLength() int32 {
	l := d.FilenameLength + d.ExtraFieldLength + d.FileCommentLength

	return int32(46 + l)
}

// ExternalFileAttributes get the attribute field from its high and low parts
func (d *ZipDirRecord) ExternalFileAttributes() uint32 {
	l := uint32(d.ExternalFileAttributesL)
	h := (uint32(d.ExternalFileAttributesH) << 16)

	return l + h
}

// RelativeOffset get the relative offset field from its high and low parts
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
	echo(" FilenameLength                     ", d.FilenameLength)
	echo(" ExtraFieldLength                   ", d.ExtraFieldLength)
	echo(" FileCommentLength                  ", d.FileCommentLength)
	echo(" DiskNumberWhereFileStarts          ", d.DiskNumberWhereFileStarts)
	echo(" InternalFileAttributes             ", d.InternalFileAttributes)
	echo(" ExternalFileAttributes()           ", d.ExternalFileAttributes())
	echo(" RelativeOffset()                   ", d.RelativeOffset())
}
