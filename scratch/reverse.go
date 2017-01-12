// Reverses a slice of ints of any size; can we do it for the slice of any type?
package main

import "fmt"

func reverse(ints []int) {
	for i, j := 0, len(ints) - 1; i < j; i, j = i + 1, j - 1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
}
func leftRotate(ints []int, d int) {
	d %= len(ints)
	reverse(ints[:d]) // 1) reverse first d
	reverse(ints[d:]) // 2) reverse remaining
	reverse(ints)     // 3) reverse all
}

func rightRotate(ints []int, d int) {
	d %= len(ints)
	reverse(ints)       // 1) reverse all
	reverse(ints[:d])   // 2) reverse first d
	reverse(ints[d:])   // 3) reverse remaining
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	reverse(ints)
	fmt.Printf("%d\n", ints) // [5 4 3 2 1]
	reverse(ints[:3])
	fmt.Printf("%d\n", ints) // [3 4 5 2 1]
	leftRotate(ints, 2)
	fmt.Printf("%d\n", ints) // [5 2 1 3 4]
	rightRotate(ints, 3)
	fmt.Printf("%d\n", ints) // [1 3 4 5 2]
}