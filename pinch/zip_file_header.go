// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

import (
	"math"
)

type ZipFileHeader struct {
	LocalFileHeaderSignature uint32
	VersionNeededToExtract   uint16
	GeneralPurposeBitFlag    uint16
	CompressionMethod        uint16
	FileLastModificationTime uint16
	FileLastModificationDate uint16
	Crc32L                   uint16
	Crc32H                   uint16
	CompressedSizeL          uint16
	CompressedSizeH          uint16
	UncompressedSizeL        uint16
	UncompressedSizeH        uint16
	FileNameLength           uint16
	ExtraFieldLength         uint16
}

func (f *ZipFileHeader) StartOffset() uint32 {
	l := uint32(f.FileNameLength) + uint32(f.ExtraFieldLength)

	return 30 + l
}

func (f *ZipFileHeader) CompressedSize() uint32 {
	return uint32(f.CompressedSizeL) + (uint32(f.CompressedSizeH) << 16)
}

func (f *ZipFileHeader) OriginalSize() uint32 {
	return uint32(f.CompressedSizeL + f.CompressedSizeH*math.MaxUint16)
}
