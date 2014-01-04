// Copyright (c) 2013-2014 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

import (
	. "github.com/smartystreets/goconvey/convey"

	"errors"
	"testing"
)

func TestPinch(t *testing.T) {
	url := "http://peterhellberg.github.io/pinch/test.zip"
	file := "data.json"

	Convey("Get", t, func() {
		Convey("example ZIP file", func() {
			json := "{\"gem\":\"pinch\",\"authors\":" +
				"[\"Peter Hellberg\",\"Edward Patel\"]," +
				"\"github_url\":\"https://github.com/peterhellberg/pinch\"}\n"

			data, _ := Get(url, file)

			So(string(data), ShouldEqual, json)
		})

		Convey("unknown ZIP file", func() {
			_, err := Get("http://example.com/unknown.zip", file)

			So(err, ShouldResemble, errors.New("404 Not Found"))
		})
	})

	Convey("GetZipDirectory", t, func() {
		Convey("example ZIP file", func() {
			entries, _ := GetZipDirectory(url)

			So(entries, ShouldResemble, map[string]ZipEntry{
				"data.json": ZipEntry{
					"data.json", 97, 114, 8, 24, 0},

				"images/pine_cone.jpg": ZipEntry{
					"images/pine_cone.jpg", 2516037, 2518084, 8, 24, 229},
			})
		})
	})
}
