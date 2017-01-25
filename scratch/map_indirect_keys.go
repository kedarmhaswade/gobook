// Sometimes we need a map or set whose keys are slices, but because a map’s keys must be comparable,
// this cannot be expressed directly. However, it can be done in two steps.
// First we define a helper function k that maps each key to a string, with the property that k(x) == k(y)
//   if and only if we consider x and y equivalent.
// Then we create a map whose keys are strings, applying the helper function to each key before we access the map.
package main

import "fmt"

func main() {
	// our mapper is the conversion function "string"
	m := make(map[string]int)
	x1 := []byte{1, 2, 3}
	x2 := []byte{2, 3, 4}
	x3 := []byte{1, 2, 3}
	m[string(x1)] = 100
	_, ok := m[string(x3)]
	if ok {
		fmt.Println("Map respects key equality function")
	} else {
		fmt.Println("OMG! Map doesn't respect key equality function")
	}
	_, ok = m[string(x2)]
	if !ok {
		fmt.Printf("There's no key: %v\n", x2)
	}
}