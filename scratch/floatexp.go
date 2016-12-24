// Finds out the value of (1 << 24) - 1
package main

import "fmt"

func main() {
	// 0 10010110 11111111111111111111111 is (1 << 24) - 1 as float32
	// +   150    --- finding this out --
	exp := 23 // 150 - 127
	// find the fractional part
	var sum float32
	var frc float32 = 1.0
	var i uint
	for i = 1; i <= 21; i++ {
		frc = float32(frc / 2.0)
		sum += frc
	}
	fmt.Printf("%d %f\n", exp, sum)
}

