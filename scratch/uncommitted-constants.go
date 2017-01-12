// Examines the uncommitted constants
package main

import "fmt"

const (
	KB = 1000
	MB = KB * KB
	GB = KB * MB
	TB = MB * MB
	PB = KB * TB
	EB = GB * GB
	ZB = KB * EB // overflows int
	YB = TB * TB // overflows int
	F100 = 212
)

func main() {
	fmt.Printf("%T %[1]d\n", YB/ZB)
	fmt.Printf("%T %[1]d\n", 212)
	fmt.Printf("%T %[1]d\n", F100)
	fmt.Printf("%T\n", 0i)
}
