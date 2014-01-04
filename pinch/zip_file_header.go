// Copyright (c) 2013-2014 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

// ZipFileHeader zip archive file header, see http://en.wikipedia.org/wiki/ZIP_(file_format)#File_headers
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

// StartOffset calculate the actual start offset for the file
func (f *ZipFileHeader) StartOffset() uint32 {
	l := uint32(f.FileNameLength) + uint32(f.ExtraFieldLength)

	return 30 + l
}

// CompressedSize get the compressed size field from its high and low parts
func (f *ZipFileHeader) CompressedSize() uint32 {
	l := uint32(f.CompressedSizeL)
	h := (uint32(f.CompressedSizeH) << 16)

	return l + h
}

// OriginalSize get the original size field from its high and low parts
func (f *ZipFileHeader) OriginalSize() uint32 {
	l := uint32(f.CompressedSizeL)
	h := (uint32(f.CompressedSizeH) << 16)

	return l + h
}
