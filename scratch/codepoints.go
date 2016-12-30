// Prints the codepoint of each character in a given String. Note that Strings are UTF-8 encoded in Go.
package main

import "fmt"

func main()  {
	var i int
	//var s = "letters: अआइउऊओऔअंअःएऐ"
	var s = "\xe4\xb8\x96\xe7\x95\x8c"
	//var s = "\u4e16\u754c"
	len := len(s)
	//for j := 0; j < len; j++ {
	//	fmt.Printf("%x\n", s[j])
	//}
	var nc int
	for i < len {
		cp, n := getCodepoint(i, s)
		fmt.Printf("decimal codepoint: %d = character: %[1]c\n", cp)
		nc++
		i += n
	}
}
// Returns the UTF-8 encoded codepoint corresponding to the byte sequence at str[pos] as a 1-, 2-, 3-, or 4-byte number
// See: https://en.wikipedia.org/wiki/UTF-8
func getCodepoint(pos int, str string) (int, int) {
	//fmt.Printf("length of string: %d\n", len(str))
	if x0 := str[pos]; x0 <= 0x7F {
		return int(x0), 1 // 0xxxxxxx (7-bit codepoint)
	} else if x1, x0 := x0, str[pos + 1];
		startsWith110(x1) && startsWith10(x0)  {
		return int(x1 & 0x1f) << 6 | int(x0 & 0x3f), 2 // 110xxxxx 10xxxxxx (11-bit codepoint)
	} else if x2, x1, x0 := x1, x0, str[pos + 2];
		startsWith1110(x2) && startsWith10(x1) && startsWith10(x0) {
		return int(x2 & 0xf) << 12 | int(x1 & 0x3f) << 6 | int(x0 & 0x3f), 3 // 1110xxxx 10xxxxxx 10xxxxxx (16-bit codepoint)
	} else if x3, x2, x1, x0 := x2, x1, x0, str[pos + 3];
		startsWith11110(x3) && startsWith10(x2) && startsWith10(x1) && startsWith10(x0) {
		return int(x3 & 0x8) << 18 | int(x2 & 0x3f) << 12 | int(x1 & 0x3f) << 6 | int(x0 & 0x3f), 4 // 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx (21-bit codepoint)
	} else {
		// panic, should I return an error?
		fmt.Printf("invalid UTF-8 byte %b at this position %d \n", str[pos], pos)
		return 0, 0
	}
}
func startsWith10(b byte) bool {
	//fmt.Printf("startswith10 %x %v\n", b, b>>6 == 2)
	return b >> 6 == 2 //10
}
func startsWith110(b byte) bool {
	//fmt.Printf("startsWith110 %x %v\n", b, b>>5 == 6)
	return b >> 5 == 6 //110
}
func startsWith1110(b byte) bool {
	//fmt.Printf("startswith1110 %x %v\n", b, b>>4 == 14)
	return b >> 4 == 14 //1110
}
func startsWith11110(b byte) bool {
	//fmt.Printf("startsWith11110 %x %v\n", b, b>>3 == 30)
	return b >> 3 == 30 // 11110
}
