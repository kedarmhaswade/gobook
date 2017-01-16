// Write a version of rotate that operates in a single pass.
// This solution uses additional O(n) memory.
package main

import "fmt"

func leftRotate(a []int, d int) []int {
	n := len(a)
	d %= n
	if d == 0 {
		return a
	}
	for i := 0; i < n; i++ {
		a = append(a, a[(i + d) % n])
	}
	return a[n:]
}
func main() {
	var a = [...]int {1, 2, 3, 4, 5}
	fmt.Printf("%d\n", leftRotate(a[:], 1))
	fmt.Printf("%d\n", leftRotate(a[:], 2))
	fmt.Printf("%d\n", leftRotate(a[:], 3))
	fmt.Printf("%d\n", leftRotate(a[:], 4))
	fmt.Printf("%d\n", leftRotate(a[:], 5))
}
