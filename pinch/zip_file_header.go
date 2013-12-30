// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

import (
	"math"
)

type ZipFileHeader struct {
	localFileHeaderSignature uint32
	versionNeededToExtract   uint16
	generalPurposeBitFlag    uint16
	compressionMethod        uint16
	fileLastModificationTime uint16
	fileLastModificationDate uint16
	crc32L                   uint16
	crc32H                   uint16
	compressedSizeL          uint16
	compressedSizeH          uint16
	uncompressedSizeL        uint16
	uncompressedSizeH        uint16
	fileNameLength           uint16
	extraFieldLength         uint16
}

func (f *ZipFileHeader) StartOffset() uint32 {
	l := uint32(f.fileNameLength) + uint32(f.extraFieldLength)

	return 30 + l
}

func (f *ZipFileHeader) CompressedSize() uint32 {
	return uint32(f.compressedSizeL) + (uint32(f.compressedSizeH) << 16)
}

func (f *ZipFileHeader) OriginalSize() uint32 {
	return uint32(f.compressedSizeL + f.compressedSizeH*math.MaxUint16)
}
