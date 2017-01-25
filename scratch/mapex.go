// Examines Maps
package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string]int)
	m["foo"] = 55
	m["bar"] = 455
	fmt.Printf("size: %d\n", len(m))
	delete(m, "foo")
	fmt.Printf("size: %d\n", len(m))
	delete(m, "baz")
	fmt.Printf("size: %d\n", len(m))
	fmt.Printf("non-existent key: xxx, value: %d\n", m["xxx"])
	m["abc"] = 424
	m["def"] = 45434
	//iterate
	for key, value := range m {
		fmt.Printf("%s : %d\n", key, value)
	}
	var names []string
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, m[name])
	}
	for k := range m {
		delete(m, k)
	}
	fmt.Printf("after deleting all keys, size of map = %d and is it same as nil? %v\n", len(m), m == nil)
	m = nil
	fmt.Printf("after explicit assignment though: is map == nil? %v\n", m == nil)
	// iterating, lookup, range and len are fine on m
	// but store is not
	// m["abc"] = 1 // panic
	key := "no-such-key"
	val, ok := m[key] // interesting, the same subscript notation returns two values in this case!
	if !ok {
		fmt.Printf("key containment test returns: %v for key: %s\n", ok, key)
		fmt.Printf("but the value returned is: %d\n", val)
	}

}

