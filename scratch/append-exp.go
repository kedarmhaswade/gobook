package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1)
	fmt.Printf("%d\n", x)
	x = append(x, x...)
	fmt.Printf("%d\n", x)
}
