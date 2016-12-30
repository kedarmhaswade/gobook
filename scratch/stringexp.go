package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Joe's body"
	fmt.Printf("%s\n", s)
	s = "\x20"
	fmt.Printf("Is \\x20 = space? : %v\n", s[0] == ' ')
	//s = "\477" // does not compile
	//fmt.Printf("%d\n", s[0])
	s = `abcd
`
	fmt.Printf("length of %s is %d\n", s, len(s))
	for _, b := range(s) {
		fmt.Printf("%d\n", b)
	}
	for _, p := range strings.Split(`a\b`, `\`) {
		fmt.Printf("%s\n", p)
	}
}
