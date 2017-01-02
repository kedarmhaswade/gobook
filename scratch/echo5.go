// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
//Echo5 is a copy of Echo4 from the book.
package main

import (
	"flag"
	"fmt"
	"strings"
	"sort"
)

func main() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var h = flag.String("h", "localhost", "host name")
	var sep = flag.String("s", " ", "separator")
	flag.Parse()
	fmt.Print(*h + ": ")
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Printf("%x %x %x %x\n", 8, 15, 31, 63)
	fmt.Printf("%x\n", 20140)
}

//!-
