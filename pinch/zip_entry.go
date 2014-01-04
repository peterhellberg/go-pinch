// Copyright (c) 2013-2014 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

// ZipEntry collected zip record fields for file 'pinching'
type ZipEntry struct {
	Filename                        string
	CompressedSize                  uint32
	UncompressedSize                uint32
	CompressionMethod               uint16
	ExtraFieldLength                uint16
	RelativeOffsetOfLocalFileHeader uint32
}

func (e *ZipEntry) echo() {
	echo("ZipEntry")
	echo(" Filename                           ", e.Filename)
	echo(" CompressedSize                     ", e.CompressedSize)
	echo(" UncompressedSize                   ", e.UncompressedSize)
	echo(" CompressionMethod                  ", e.CompressionMethod)
	echo(" ExtraFieldLength                   ", e.ExtraFieldLength)
	echo(" RelativeOffsetOfLocalFileHeader    ", e.RelativeOffsetOfLocalFileHeader)
}
