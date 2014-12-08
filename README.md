# Go Pinch

[![Build Status](https://travis-ci.org/peterhellberg/go-pinch.png?branch=master)](https://travis-ci.org/peterhellberg/go-pinch)
[![GoDoc](https://godoc.org/github.com/peterhellberg/go-pinch/pinch?status.png)](https://godoc.org/github.com/peterhellberg/go-pinch/pinch)

## Retrieve a file from inside a zip archive, over the network!

Pinch makes it possible to download a specific file from within
a ZIP archive over HTTP/1.1, using nothing but the Go Standard
Library ([net/http](http://golang.org/pkg/net/http/) and
[compress/flate](http://golang.org/pkg/compress/flate/))

Earlier versions were written in [Objective-C](https://github.com/epatel/pinch-objc), [Ruby](https://github.com/peterhellberg/pinch) and [Java](https://github.com/carlbenson/Pinch)

*STATUS:* Working, but not in any known use :)

## Installation

```bash
go get -u github.com/peterhellberg/go-pinch
```

## Usage

```bash
$ go-pinch http://example/path/to.zip file.json
```

Or from Go directly:

```go
package main

import (
	"github.com/peterhellberg/go-pinch/pinch"
	"os"
)

func main() {
	file, _ := pinch.Get("http://peterhellberg.github.com/pinch/test.zip", "data.json")

	os.Stdout.Write(file)
}
```
