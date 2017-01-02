// Write const declarations for KB, MB, up through YB as compactly as you can.
package main

import (
	"fmt"
	"strconv"
	"exercises/ch3/format"
)

const (
	X float32 = 9
	KB = 1000
	MB = KB * KB
	GB = KB * MB
	TB = MB * MB
	PB = KB * TB
	EB = GB * GB
	ZB = KB * EB // overflows int
	YB = TB * TB // overflows int
)

func main() {
	//fmt.Printf("%d %d %d %d %d %d %d %d\n", KB, MB, GB, TB, PB, EB, ZB, YB) // compile error
	fmt.Printf("%d %d %d %d %d %d\n", KB, MB, GB, TB, PB, EB)
	fmt.Printf("%T %T %T %T %T %T\n", KB, MB, GB, TB, PB, EB)
	fmt.Printf("%T %[1]d\n", YB/ZB)
	fmt.Printf("%s %s %s %s %s %s\n", format.CommaFloat(strconv.Itoa(KB), false), format.CommaFloat(strconv.Itoa(MB), false),
		format.CommaFloat(strconv.Itoa(GB), false), format.CommaFloat(strconv.Itoa(TB), true),
		format.CommaFloat(strconv.Itoa(PB), true), format.CommaFloat(strconv.Itoa(EB), true))
}
