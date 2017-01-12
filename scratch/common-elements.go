// Given two slices, prints the elements that are common to them
package main

import "fmt"

func main() {
	a := [...]int {2, 1, 4, 2, 1}
	printCommonSlow(a[0:2], a[2:4])
	printCommonSlow(a[0:2], a[4:])
}
func printCommonSlow(a []int, b []int) {
	fmt.Println("Printing elements that are common to two slices")
	for _, e1 := range a {
		for _, e2 := range b {
			if e1 == e2 {
				fmt.Printf("%v\n", e1)
			}
		}
	}
}
