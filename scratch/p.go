package main

import (
	"fmt"
	"math"
)

func main() {
	var v int = 3
	var p *int = &v
	*p++
	fmt.Printf("%d\n", *p)
	fmt.Printf("%d\n", v)
	delta(3, 5)
	notEscaping()
	plusPlus()
}
func delta(old, new int) {
	fmt.Printf("%d\n", int(math.Abs(float64(new-old))))
}

func notEscaping() {
	y := new(int)
	*y = 1
	fmt.Printf("%v, %v\n", y, &*y) // 0xc82000a308, 0xc82000a308
}
func plusPlus() {
	var f float64 = 1.2
	f++
	fmt.Printf("%v\n", f)
}
