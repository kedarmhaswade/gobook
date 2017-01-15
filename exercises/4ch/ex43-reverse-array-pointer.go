// Reverses an array by taking a pointer to it instead of a slice from it
package main

import (
	"fmt"
)

func reverse(ptr *[4]int)  {
	for i, j := 0, len(*ptr) - 1; i < j; i, j = i + 1, j - 1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}
func main() {
	a := [...]int {1, 2, 3, 10}
	reverse(&a)
	fmt.Printf("%d\n", a)
}
