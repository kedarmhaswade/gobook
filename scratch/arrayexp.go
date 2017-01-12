package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3}
	b := []int{1, 2, 3, 4}
	fmt.Printf("%d\n", len(a)) // 3
	fmt.Printf("%d\n", len(b)) // 4
	b = a
	fmt.Printf("%d\n", len(b)) // 3
	fmt.Printf("%T %T\n", a, b) // 3
	ab := byteProvider()
	fmt.Printf("%T \n", ab)
	fmt.Printf("%b\n%b\n", 120, 88)
}
func byteProvider() []byte {
	return []byte{100, 200}
}
