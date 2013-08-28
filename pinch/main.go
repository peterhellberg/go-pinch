package main

import (
	"github.com/peterhellberg/go-pinch"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	pinch.Get(getArgs(os.Args))
}

func getArgs(args []string) (string, string) {
	// Make sure that we got three command line arguments
	if len(args) != 3 {
		log.Fatalln("Usage: pinch <url> <file>")
	}

	// Parse the URI parameter
	_, err := url.ParseRequestURI(args[1])

	if err != nil || !strings.HasPrefix(args[1], "http") {
		log.Fatalln("Invalid URL")
	}

	// Return the two arguments
	return args[1], args[2]
}
