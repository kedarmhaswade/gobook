// Examines types
package main

import (
	"time"
	"fmt"
)

func main() {
	var duration time.Duration
	var m5 = time.Minute * 5
	fmt.Printf("%T\n", duration)
	fmt.Printf("%T %[1]v\n", m5)
	fmt.Printf("%T\n", m5.Hours())
}

