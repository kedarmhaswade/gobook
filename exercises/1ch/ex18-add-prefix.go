// Exercise 1.8: Modify fetch to add the prefix http://
// to each argument URL if it is missing. You might want to use strings.HasPrefix.
package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"strings"
)

func main()  {
	prefix := "http://"
	for _, url := range os.Args[1:] {
		if ! strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex18: %v\n", err)
			os.Exit(1)
		}
		length, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex18: unable to copy to stdout, code: %v\n", err)
			return
		}
		fmt.Printf("Downloaded %d bytes\n", length)
	}
}
