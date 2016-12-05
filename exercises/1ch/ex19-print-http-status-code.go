// Exercise 1.9: Modify fetch to also print the HTTP status code, found in resp.Status.
package main

import (
	"os"
	"strings"
	"net/http"
	"fmt"
	"io"
)

func main()  {
	prefix := "http://"
	for _, url := range os.Args[1:] {
		if ! strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex19-print-http-status-code.go: error getting %s, code: %v\n", url, err)
			continue
		}
		length, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("ex19-print-http-status-code.go: error reading URL: %v\n", err)
			continue
		}
		fmt.Printf("HTTP status code: %v\n", resp.Status)
		fmt.Printf("Downloaded %d bytes\n", length);
	}

}