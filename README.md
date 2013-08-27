# Go Pinch

*TODO:* Actually write the code.

Retrieve a file from inside a zip file, over the network!

Pinch makes it possible to download a specific file from within 
a ZIP file over HTTP/1.1, using nothing but the Go Standard 
Library ([net/http](http://golang.org/pkg/net/http/) and 
[compress/zlib](http://golang.org/pkg/compress/zlib/))

The first version was written in Objective-C and we thought it 
would be cool if we could bring that functionality to Ruby, 
I’m now about to translate the code to Go :)

## Usage

```bash
$ pinch http://example/path/to.zip file.json
```

Or from Go directly:

```go
package main

import (
  "github.com/peterhellberg/go-pinch"
  "net/url"
  )

func main() {
  url, _ = uri.ParseRequestURI("http://example/path/to.zip")
  
  pinch.Get(url, "file.json")
}
```
