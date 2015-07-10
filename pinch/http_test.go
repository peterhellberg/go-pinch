package pinch

import (
	. "github.com/smartystreets/goconvey/convey"

	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	url := "http://peterhellberg.github.io/pinch/test.zip"

	Convey("getContentLength", t, func() {
		Convey("correct length for example", func() {
			length, _ := getContentLength(url)

			So(length, ShouldEqual, 2516612)
		})

		Convey("length 0 if HTTP status != 200", func() {
			length, _ := getContentLength("http://example.com/missing.zip")

			So(length, ShouldEqual, 0)
		})
	})

	Convey("fetchPartialData", t, func() {
		Convey("fetch a range of bytes", func() {
			bytes, _ := fetchPartialData(url, 1, 11)

			So(bytes, ShouldResemble, []byte{
				75, 3, 4, 20, 0, 0, 0, 8, 0, 54, 130})
		})
	})

	Convey("rangeHTTPClient", t, func() {
		Convey("creates HTTP client from bytesRange", func() {
			bytesRange := "bytes=2=22"

			req, _ := http.NewRequest("GET", "http://example.com", nil)

			client := rangeHTTPClient(bytesRange)

			So(client.CheckRedirect(req, nil), ShouldBeNil)
			So(req.Header.Get("Range"), ShouldEqual, bytesRange)
		})
	})
}
