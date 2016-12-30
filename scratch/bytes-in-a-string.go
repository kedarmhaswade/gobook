// Prints the values of bytes in a given string
package main

import "fmt"

func main() {
	s := "₹"
	len := len(s)
	fmt.Printf("number of bytes in of s = %d\n", len)
	for i := 0; i < len; i++ {
		fmt.Printf("byte%d: %x\n", i, s[i])
	}
}
