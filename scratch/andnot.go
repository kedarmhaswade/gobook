package main

import "fmt"

func main() {
	var x uint8 = 1
	var y uint8 = 2
	var ny uint8 = ^y
	fmt.Printf("%d\n", x &^ y) // the difference operator
	fmt.Printf("%d\n", x & ny)
}
