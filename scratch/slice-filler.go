// Fills the given slice of strings
package main

import "fmt"

func fill(ss []string) {
	ss[1] = "foo" // panics!
}
func main() {
	ss  := []string{"first"}
	fill(ss)
	fmt.Printf("%v\n", ss)
}
