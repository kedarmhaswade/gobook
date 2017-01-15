// Uses append function to append an int to an int array
package main

import (
	"fmt"
	"math/rand"
)

//import "fmt"

// appends the given int to the given slice and returns the new slice; does not use append library function
func appendIntBook(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen] // this is a new slice!
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
func appendInt(x []int, y int) []int {
	lenx := len(x)
	if cap(x) > lenx {
		//fmt.Printf("%d has enough capacity: %d\n", x, cap(x))
		x := x[:lenx + 1]
		x[lenx] = y
		return x
	}
	newCap := cap(x) * 2
	if newCap == 0 { // special case when cap(x) is 0
		newCap = 1
	}
	z := make([]int, lenx+1, newCap)
	fmt.Printf("resized to %d\n", newCap)
	copy(z, x)
	z[lenx] = y
	return z
}

func main() {
	var a = []int(nil)
	var x = a[:]
	for i := 0; i < 10; i++ {
		x = appendInt(x, rand.Intn(30))
		fmt.Printf("%d\n", x)
	}

}
