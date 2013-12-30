// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

type ZipEntry struct {
	Filename                        string
	CompressedSize                  uint32
	UncompressedSize                uint32
	CompressionMethod               uint16
	ExtraFieldLength                uint16
	RelativeOffsetOfLocalFileHeader uint32
}
