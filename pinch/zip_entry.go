// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

type ZipEntry struct {
	Filename                        string
	compressedSize                  uint32
	uncompressedSize                uint32
	compressionMethod               uint16
	extraFieldLength                uint16
	relativeOffsetOfLocalFileHeader uint32
}
