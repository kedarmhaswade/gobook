// Calls various PopCount functions
package main

import (
	"exercises/ch2/popcount"
	//"math/rand"
	"fmt"
)

func main() {
	// Make sure all the ways agree
	limit := 1 << 26
	for i := 1; i < limit; i++ {
		arg := uint64(i)
		n1 := popcount.PopCount(arg)
		n2 := popcount.PopCountLoop(arg)
		n3 := popcount.PopCountBitwise(arg)
		if n1 != n2 || n1 != n3 {
			fmt.Printf("This is a bug for n = %d, n1 = %d, n2 = %d, n3 = %d\n", i, n1, n2, n3)
			return
		}
	}
	fmt.Printf("All the ways of getting population count seem to agree with %d integers that we tested", limit)
}