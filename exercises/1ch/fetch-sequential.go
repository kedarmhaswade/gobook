// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetch-sequential fetches URLs in sequentially and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		fetchSequentially(url) // fetch them sequentially
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchSequentially(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error fetching %s, code: %v\n", url, err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		fmt.Printf("error while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs  %7d  %s\n", secs, nbytes, url)
}

//!-
