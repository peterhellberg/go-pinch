package main

import (
	"github.com/peterhellberg/go-pinch"
	"log"
	"net/url"
	"os"
)

func main() {
	url, fn = getArgs(os.Args)

	pinch.Get(url, fn)
}

func getArgs(args []string) (*url.URL, string) {
	// Make sure that we got three command line arguments
	if len(args) != 3 {
		log.Fatalln("Usage: pinch <url> <file>")
	}

	// Parse the URI parameter
	var url, err = url.ParseRequestURI(args[1])
	if err != nil {
		log.Fatalln("Invalid URL")
	}

	return url, args[2]
}
