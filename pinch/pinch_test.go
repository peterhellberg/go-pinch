package pinch

import (
	. "github.com/smartystreets/goconvey/convey"

	"errors"
	"testing"
)

func TestPinch(t *testing.T) {
	url := "http://peterhellberg.github.io/pinch/test.zip"
	file := "data.json"

	exampleData := []byte(
		"{\"gem\":\"pinch\",\"authors\":[\"Peter Hellberg\",\"Edward Patel\"]," +
			"\"github_url\":\"https://github.com/peterhellberg/pinch\"}\n")

	exampleEntries := map[string]ZipEntry{
		"data.json": ZipEntry{
			"data.json", 97, 114, 8, 24, 0},

		"images/pine_cone.jpg": ZipEntry{
			"images/pine_cone.jpg", 2516037, 2518084, 8, 24, 229},
	}

	Convey("Get", t, func() {
		Convey("known file from example ZIP archive", func() {
			data, _ := Get(url, file)

			So(data, ShouldResemble, exampleData)
		})

		Convey("missing file from example ZIP archive", func() {
			_, err := Get(url, "missing.zip")

			So(err, ShouldResemble, errors.New("file not found in archive"))
		})

		Convey("missing ZIP archive", func() {
			_, err := Get("http://example.com/missing.zip", file)

			So(err, ShouldResemble, errors.New("404 Not Found"))
		})
	})

	Convey("GetZipFile", t, func() {
		Convey("file from URL and ZipEntry", func() {
			data, _ := GetZipFile(url, exampleEntries["data.json"])

			So(data, ShouldResemble, exampleData)
		})
	})

	Convey("GetZipDirectory", t, func() {
		Convey("example ZIP file", func() {
			entries, _ := GetZipDirectory(url)

			So(entries, ShouldResemble, exampleEntries)
		})
	})
}
