// Examines a few fmt package tricks
package main

import (
	"fmt"
	"math"
)

func main() {
	// integers
	x := 32
	fmt.Printf("%d =  b%[1]b = O%[1]o = x%[1]x = X%[1]x\n", x)
	// runes: note a Rune is a synonym
	ascii := 'a'
	unicode := '™'
	newline := '\n'
	fmt.Printf("%d %[1]x %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "8482 ™ '™'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
	fmt.Printf("%d\n", math.MaxInt8)
}
