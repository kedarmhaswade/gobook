package main

import (
	"fmt"
	"math"
)

func main() {
	//var f float32 = 16777215
	fmt.Printf("%032b\n", math.Float32bits(16777215))
	fmt.Printf("%032b\n", math.Float32bits(16777216))
	fmt.Printf("%032b\n", math.Float32bits(16777217))
	fmt.Printf("%032b\n", math.Float32bits(16777218))
	fmt.Printf("%032b\n", math.Float32bits(16777219))
	fmt.Printf("%032b\n", math.Float32bits(16777220))
}
