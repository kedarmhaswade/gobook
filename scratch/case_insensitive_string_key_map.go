// Demonstrates comparable key types when you want a definition of equality other than ==,
// such as case-insensitive comparisons for strings
package main

import (
	"strings"
	"fmt"
)

func k(key string) string {
	return strings.ToLower(key)
}
func main() {
	m := map[string]int{}
	m[k("abc")] = 1
	name := "Abc"
	_, ok := m[k(name)]
	if ok {
		fmt.Printf("name %v exists!\n", name)
	} else {
		fmt.Printf("name %v does not exist\n", name)
	}
	name = "deF"
	_, ok = m[k(name)]
	if ok {
		fmt.Printf("name %v exists!\n", name)
	} else {
		fmt.Printf("name %v does not exist\n", name)
	}
}
