/*---------------------------------------------------------------------------
 
 Copyright (c) 2013 Peter Hellberg, Edward Patel
 
 Permission is hereby granted, free of charge, to any person obtaining
 a copy of this software and associated documentation files (the
 "Software"), to deal in the Software without restriction, including
 without limitation the rights to use, copy, modify, merge, publish,
 distribute, sublicense, and/or sell copies of the Software, and to
 permit persons to whom the Software is furnished to do so, subject to
 the following conditions:

 The above copyright notice and this permission notice shall be
 included in all copies or substantial portions of the Software.

 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
 LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
 OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 
 ---------------------------------------------------------------------------*/

package main

import (
    "go-pinch"
    "os"
    "fmt"
    "log"
    "net/url"
    "strings"
)

func main() {

    args := getArgs(os.Args)

    url := args[0]

    entries, err := pinch.GetZipDirectory(url)
    if err != nil {
        fmt.Println(err)
        log.Fatalln("exiting")
    }

    if len(args) == 2 {
        
        var entry pinch.ZipEntry = entries[args[1]]

        if len(entry.Filename) > 0 {

            file, err := pinch.GetZipFile(url, entry)
            if err != nil {
                fmt.Println(err)
                log.Fatalln("exiting")
            }

            os.Stdout.Write(file)

        } else {

            fmt.Printf("File not found\n")
        
        }
    
    } else {

        for _, entry := range entries {
            fmt.Println(entry.Filename)
        }

    }
}

func getArgs(args []string) []string {

    // Make sure that we got three command line arguments
    if len(args) < 2 || len(args) > 3 {
        fmt.Println("Usage: pinch <url> [ <file> ]")
        log.Fatalln("exiting")
    }

    // Parse the URI parameter
    _, err := url.ParseRequestURI(args[1])

    if err != nil || !strings.HasPrefix(args[1], "http") {
        fmt.Println("Invalid URL")
        log.Fatalln("exiting")
    }

    if len(args) == 2 {
        return []string{ args[1] }
    }

    return []string{ args[1], args[2] }
}
