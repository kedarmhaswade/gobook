// Examines slice comparability
package main

import "fmt"

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	c := []byte{'o', 'l', 'a'}
	s := b[1:4]  // sharing the same storage as b
	//fmt.Printf("%v\n", c == s) // note: slices are not comparable using the == sign, you have to write it
	for i := range c {
		if i >= len(s) || s[i] != c[i] {
			fmt.Println("slices are unequal")
			break
		}
	}
	fmt.Println("slices are equal!")
}
