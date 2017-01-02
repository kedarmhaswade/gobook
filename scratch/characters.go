// Prints the characters in a given string
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "\xe4\xb8\x96\xe7\x95\x8c"
	for _, c := range s {
		fmt.Printf("%c %[1]U\n", c)
	}
	fmt.Println(len(s))                    // "6"
	fmt.Println(utf8.RuneCountInString(s)) // "2"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}
