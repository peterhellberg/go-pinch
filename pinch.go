package pinch

import (
	_ "compress/zlib"
	"fmt"
	"net/http"
)

func Get(url string, fn string) {
	of := int64(0)
	cl := getContentLength(url)

	if cl >= 4096 {
		of = cl - 4096
	}

	fmt.Println("URL            ", url)
	fmt.Println("File name      ", fn)
	fmt.Println("Content-Length ", cl)
	fmt.Println("Offset         ", of)
}

func getContentLength(url string) int64 {
	resp, err := http.Head(url)
	if err != nil {
		fmt.Println(err)
	}
	return resp.ContentLength
}

func getData(sof int64, eof int64) {
}
