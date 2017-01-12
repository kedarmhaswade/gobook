// Demonstrates that arrays are passed by value in Go
package main

import "fmt"

func main() {
	ints := [...]int{1, 4, 12}
	square1(ints)            // the argument passed is a true copy of the ints array
	fmt.Printf("%d\n", ints) // ints is not affected!
	square2(&ints)           // the argument passed is a reference to the ints array
	fmt.Printf("%d\n", ints) // ints is affected!
}
func square1(ints [3]int) {
	for i, x := range ints {
		ints[i] = x * x
	}
}
func square2(ptrToInts *[3]int) {
	for i, x := range ptrToInts {
		ptrToInts[i] = x * x
	}
}
