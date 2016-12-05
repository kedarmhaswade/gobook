// Exercise 1.7: The function call io.Copy(dst, src) reads from src and writes to dst. 
// Use it instead of ioutil.ReadAll to copy the response body to os.Stdout without 
// requiring a buffer large enough to hold the entire stream. 
// Be sure to check the error result of io.Copy.
package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

func main()  {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		length, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: unable to copy to stdout, code: %v\n", os.Args[0], err)
			return
		}
		fmt.Printf("%s: Downloaded %d bytes\n", os.Args[0], length)
	}
}
