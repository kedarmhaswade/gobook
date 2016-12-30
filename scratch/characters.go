// Prints the characters in a given string
package main

import "fmt"

func main() {
	s := "\xe4\xb8\x96\xe7\x95\x8c"
	for _, c := range s {
		fmt.Printf("%c %[1]U\n", c)
	}
}
