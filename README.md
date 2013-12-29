# Go Pinch

*TODO:* Clean up the code.

Retrieve a file from inside a zip file, over the network!

Pinch makes it possible to download a specific file from within 
a ZIP file over HTTP/1.1, using nothing but the Go Standard 
Library ([net/http](http://golang.org/pkg/net/http/) and 
[compress/flate](http://golang.org/pkg/compress/flate/))

Earlier versions was written in [Objective-C](https://github.com/epatel/pinch-objc) and [Ruby](https://github.com/peterhellberg/pinch)

## Usage

```bash
$ ./pinch http://example/path/to.zip file.json
```

Or from Go directly:

```go
package main

import (
  "go-pinch"
)

func main() {
  
  entries, err := pinch.GetZipDirectory("http://example/path/to.zip")

  file, err := pinch.GetZipFile(url, entries["file.json"])

}
```
