// Exercise 1.2: Modify the echo program to print the index and value
// of each of its command line arguments, one per line.
package main

import (
	"os"
	"fmt"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d\t%s\n", i, arg)
	}
}
