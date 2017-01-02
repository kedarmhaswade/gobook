// Explores enums
package main

import "fmt"

type Color int

const  (
	RED = 1 << iota
	GREEN
	BLUE
)

func main() {
	c := RED
	fmt.Printf("%v\n", c)
	c = GREEN
	fmt.Printf("%v\n", c)
	c = BLUE
	fmt.Printf("%v\n", c)
	fmt.Printf("%T\n", c)
}