package main

import (
	. "github.com/smartystreets/goconvey/convey"

	"os/exec"
	"testing"
)

func TestPinchCommand(t *testing.T) {
	Convey("./pinch-test", t, func() {
		url := "http://peterhellberg.github.io/pinch/test.zip"

		exampleData := []byte(
			"{\"gem\":\"pinch\",\"authors\":[\"Peter Hellberg\",\"Edward Patel\"]," +
				"\"github_url\":\"https://github.com/peterhellberg/pinch\"}\n")

		Convey("with invalid URL", func() {
			out, err := runPinch("foo", "")

			So(err, ShouldNotBeNil)
			So(out, ShouldContainSubstring, "Invalid URL")
		})

		Convey("with example URL", func() {
			out, _ := runPinch(url, "")

			So(out, ShouldContainSubstring, "data.json")
			So(out, ShouldContainSubstring, "images/pine_cone.jpg")
		})

		Convey("with example URL and file name", func() {
			out, _ := runPinch(url, "data.json")

			So(out, ShouldResemble, string(exampleData))
		})
	})
}

func runPinch(url string, fn string) (string, error) {
	out, err := exec.Command("go", "run", "main.go", url, fn).CombinedOutput()

	return string(out), err
}
