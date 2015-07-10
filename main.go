/*

Retrieve a file from inside a zip file, over the network!

Pinch makes it possible to download a specific file from within
a ZIP file over HTTP/1.1, using nothing but the Go Standard Library.

(http://golang.org/pkg/net/http/ and http://golang.org/pkg/compress/flate/)

Earlier versions were written in Objective-C and Ruby.

Installation

Just go get it:

		go get -u github.com/peterhellberg/go-pinch

Usage

You get a list of files if you only pass a URL:

		go-pinch http://example/path/to.zip

In order to pinch one of the files in the ZIP:

		go-pinch http://example/path/to.zip file.json > file.json

*/
package main

import (
	"log"
	"net/url"
	"os"

	"github.com/peterhellberg/go-pinch/pinch"
)

func main() {
	url, fn := getArgs(os.Args)

	if len(fn) > 0 {
		writeFileToStdout(url, fn)
	} else {
		listFilesToStdout(url)
	}
}

func getArgs(args []string) (string, string) {
	// Make sure that we got two or three command line arguments
	if len(args) < 2 || len(args) > 3 {
		fatal("Usage: go-pinch <url> [file]")
	}

	// Parse the URI parameter
	_, err := url.ParseRequestURI(args[1])

	if err != nil || args[1][0:4] != "http" {
		fatal("Invalid URL")
	}

	if len(args) == 2 {
		return args[1], ""
	}

	return args[1], args[2]
}

func writeFileToStdout(url, fn string) {
	file, err := pinch.Get(url, fn)

	handleError(err)

	os.Stdout.Write(file)
}

func listFilesToStdout(url string) {
	entries, err := pinch.GetZipDirectory(url)

	handleError(err)

	for _, entry := range entries {
		os.Stdout.Write([]byte(entry.Filename + "\n"))
	}
}

func fatal(v ...interface{}) {
	log.Fatalln(v...)
}

func handleError(err error) {
	if err != nil {
		fatal(err)
	}
}
