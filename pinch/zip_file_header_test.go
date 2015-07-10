package pinch

import (
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestZipFileHeader(t *testing.T) {
	f := ZipFileHeader{
		uint32(1),  // LocalFileHeaderSignature
		uint16(2),  // VersionNeededToExtract
		uint16(3),  // GeneralPurposeBitFlag
		uint16(8),  // CompressionMethod
		uint16(4),  // FileLastModificationTime
		uint16(5),  // FileLastModificationDate
		uint16(6),  // Crc32L
		uint16(7),  // Crc32H
		uint16(12), // CompressedSizeL
		uint16(2),  // CompressedSizeH
		uint16(11), // UncompressedSizeL
		uint16(13), // UncompressedSizeH
		uint16(24), // FilenameLength
		uint16(48)} // ExtraFieldLength

	Convey("Fields", t, func() {
		So(f.LocalFileHeaderSignature, ShouldEqual, 1)
		So(f.VersionNeededToExtract, ShouldEqual, 2)
		So(f.GeneralPurposeBitFlag, ShouldEqual, 3)
		So(f.CompressionMethod, ShouldEqual, 8)
		So(f.FileLastModificationTime, ShouldEqual, 4)
		So(f.FileLastModificationDate, ShouldEqual, 5)
		So(f.Crc32L, ShouldEqual, 6)
		So(f.Crc32H, ShouldEqual, 7)
		So(f.CompressedSizeL, ShouldEqual, 12)
		So(f.CompressedSizeH, ShouldEqual, 2)
		So(f.UncompressedSizeL, ShouldEqual, 11)
		So(f.UncompressedSizeH, ShouldEqual, 13)
		So(f.FilenameLength, ShouldEqual, 24)
		So(f.ExtraFieldLength, ShouldEqual, 48)
	})

	Convey("Methods", t, func() {
		Convey("StartOffset()", func() {
			So(f.StartOffset(), ShouldEqual, 30+48+24)
		})

		Convey("CompressedSize()", func() {
			So(f.CompressedSize(), ShouldEqual, 131084)
		})

		Convey("OriginalSize()", func() {
			So(f.OriginalSize(), ShouldEqual, 851979)
		})
	})
}
