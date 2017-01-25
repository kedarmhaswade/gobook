// Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace)
// in a UTF-8-encoded []byte slice into a single ASCII space.
package main

import (
	"unicode"
	"fmt"
	"unicode/utf8"
)

func squash(x [] byte) []byte {
	n := len(x)
	prev := x[0]
	j := 1
	for i := 1; i < n; i++ {
		curr := x[i]
		prevSpace := unicode.IsSpace(rune(prev))
		currSpace := unicode.IsSpace(rune(curr))
		if ! (prevSpace && currSpace) {
			x[j] = x[i]
			j++
			prev = curr
		}
	}
	return x[:j]
}
func main() {
	x := squash([]byte("      मेदार  जीपा a   b c d          e f g \u4e16           \u754c   "))
	fmt.Printf("%s\n", string(x))
	x = []byte("x")
	fmt.Printf("%s\n", string(x))
	for _, c := range "name: म्हसवडे केदार" {
		fmt.Printf("%c %d\n", c, utf8.RuneLen(c))
	}
}
