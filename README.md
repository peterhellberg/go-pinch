# Go Pinch

[![Build Status](https://travis-ci.org/peterhellberg/go-pinch.png?branch=master)](https://travis-ci.org/peterhellberg/go-pinch)
[![GoDoc](https://godoc.org/github.com/peterhellberg/go-pinch/pinch?status.png)](https://godoc.org/github.com/peterhellberg/go-pinch/pinch)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/go-pinch#license-mit)

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

## License (MIT)

Copyright (c) 2013-2015 [Peter Hellberg](http://c7.se/), [Edward Patel](http://memention.com/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
