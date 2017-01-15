// Can arrays or slices be compared? Only to nil.
package main

import "fmt"

func main() {
	//a := [3]int {}
	b := []int {}
	c := []int {}
	//fmt.Printf("is b == c? %b", b == c) // slice can only be compared to nil
	b = nil
	fmt.Printf("is b == nil? %v\n", b == nil) // slice can only be compared to nil
	fmt.Printf("is c == nil? %v\n", c == nil) // slice can only be compared to nil
}
