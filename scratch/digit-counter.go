// Given a string, count the number of digits in it
package main

import (
	"fmt"
)

func countDigits(s string) int {
	cnt := 0
	for _, x := range s {
		if x >= '0' && x <= '9' {
			cnt ++
		}
	}
	return cnt
}
func main() {
	fmt.Printf("%d\n", countDigits("abc123"))
	fmt.Printf("%d\n", countDigits("\u4e161\u754c2"))
	fmt.Printf("%d\n", countDigits("abbbadadasda"))
}
