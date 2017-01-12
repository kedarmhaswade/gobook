// Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag
// to print the SHA384 or SHA512 hash instead.
package main

import (
	"os"
	"strconv"
	"fmt"
)

func main() {
	// how to grow slice?
	nargs := len(os.Args)
	dlen := 256 // length of the digest
	if nargs == 2 {
		arg := os.Args[1]
		d, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("invalid digest length: %s, exiting ...\n", arg)
			return
		}
		dlen = d
	}
	fmt.Printf("%d\n", dlen)
}
