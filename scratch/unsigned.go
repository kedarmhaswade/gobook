package main

import "fmt"

func main() {
	var x uint8 = 0;
	x--
	fmt.Printf("%d\n", x) // never becomes negative
	x++
	fmt.Printf("%d\n", x) // never becomes negative
}
