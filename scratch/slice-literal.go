// A slice literal example
package main

import "fmt"

func main() {
	const f int = 12
	s := []string {0: "a", 2: "b", f: "f"}
	fmt.Printf("%d\n", len(s))
	fmt.Printf("%s\n", s)
}
