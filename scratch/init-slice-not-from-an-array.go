// Is it possible to declare slice by itself?
package main

import "fmt"

func main() {
	//var ints []int
	ints := []int{}
	fmt.Printf("%d\n", cap(ints))
}