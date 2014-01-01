// Copyright (c) 2013 Peter Hellberg, Edward Patel.
// Licensed under the MIT License found in the LICENSE file.

package pinch

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getContentLength(url string) (int64, error) {
	resp, err := http.Head(url)

	if err != nil {
		return 0, err
	}

	if resp.StatusCode != 200 {
		return 0, errors.New(resp.Status)
	}

	return resp.ContentLength, nil
}

func fetchPartialData(url string, sof, eof int64) ([]byte, error) {
	echo("Pinch", sof, "-", eof, "(", eof-sof, ")")

	// Create the bytes range string
	bytesRange := fmt.Sprintf("bytes=%d-%d", sof, eof)

	// Create a request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	// Add the "Range:" header to the initial request
	req.Header.Add("Range", bytesRange)

	// Create a client to be used to add the "Range:" header
	client := rangeHttpClient(bytesRange)

	// Now run the request
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// Clean up
	defer resp.Body.Close()

	// Get body
	body, err := ioutil.ReadAll(resp.Body)

	// And return it...
	return body, err
}

func rangeHttpClient(bytesRange string) *http.Client {
	return &http.Client{
		// Go net/http does not keep headers in redirects, add specifically
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return errors.New("To many redirects")
			}

			req.Header.Add("Range", bytesRange)

			return nil
		},
	}
}
