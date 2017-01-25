// Prints the values of bytes in a given string
package main

import "fmt"

func main() {
	s := "₹"
	writeBytes(s)
	s = "म"
	writeBytes(s)
	s = "्"
	writeBytes(s)
	s = "म्ह"
	writeBytes(s)
	s = "ह"
	writeBytes(s)
	s = "मह"
	writeBytes(s)
	s = "क"
	writeBytes(s)
	s = "a"
	writeBytes(s)
	s = "\u00a2"
	writeBytes(s)
}
func writeBytes(s string) {
	n := len(s)
	fmt.Printf("number of bytes in %s = %d\n", s, n)
	for i := 0; i < n; i++ {
		fmt.Printf("byte%d: %x\n", i, s[i])
	}
}
