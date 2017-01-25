// Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string,
// in place. Can you do it without allocating new memory?
package main

import (
	"unicode/utf8"
	"fmt"
)

// Accepts a byte slice that represents UTF-8 encoded characters and reverses those characters in place.
// Mutates the given slice.
func reverseUtf8Chars(s []byte)  {
	//fmt.Printf("%x\n", s)
	// We play only with bytes, so temporarily the []byte will be invalid, but when the method returns,
	// it is guaranteed that it is a valid UTF-8 encoding (if the given []byte represents valid UTF-8).
	i := 0  // the reading index
	n := len(s)
	for i < n {
		r, bc := utf8.DecodeRune(s[i:]) //character and number of bytes per character
		if r == utf8.RuneError {
			fmt.Printf("Invalid UTF-8 byte sequence at index %d\n", i)
			// don't continue
			return
		}
		reverseSlice(s, i, bc)
		//fmt.Printf("loop: %x\n", s)
		i += bc
	}
	//fmt.Printf("%x\n", s)
	reverseSlice(s, 0, n)
}
// reverses in place, the nb bytes of the slice p starting at index begInc (inclusive)
func reverseSlice(p []byte, begInc int, nb int) {
	endExc := begInc + nb // exclusive
	for i, j := begInc, endExc - 1; i < j; i, j = i + 1, j - 1 {
		p[i], p[j] = p[j], p[i]
	}
}
func main() {
	s := "n: क"
	call(s)
	s = "केदार"
	call(s)
	s = "दीपा"
	call(s)
	s = "अपूर्व"
	call(s)
	s = "ऋजुता"
	call(s)
	s = "माझेानाव: KEDAR"
	call(s)
	s = "राम"
	call(s)
	s = "राज"
	call(s)
	s = "आजार"
	call(s)
}
func call(s string) {
	b := []byte(s)
	reverseUtf8Chars(b)
	fmt.Printf("%s\n", string(b))
}