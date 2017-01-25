// Examines if slices are values (and not pointers)
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	s1 := a[0:3]
	s2 := a[2:4]
	fmt.Printf("%p\n", s1)
	fmt.Printf("%d\n", cap(s1)) // this would be 5
	fmt.Printf("%p\n", s2)
	fmt.Printf("%d\n", cap(s2)) // this would be 3, not 5
	s2[0] = 44
	fmt.Printf("%d\n", s1) // affects s1
}