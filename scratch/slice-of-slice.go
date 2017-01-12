// What is the capacity of the slice of a slice?
package main

import "fmt"

func main() {
	var master [100]int
	var s1 = master[:50]
	var s11 = s1[:25]
	fmt.Printf("len: %d, cap: %d\n", len(s11), cap(s11))
}
