// Write an in-place function to eliminate adjacent duplicates in a []string slice.
package main

import "fmt"

func removeAdjacentDupStringsInPlace(strings []string) []string {
	if  len(strings) <= 1 {
		return strings
	}
	var prev string = strings[0]
	j := 1 // insert index
	for i := 1; i < len(strings); i++ {
		curr := strings[i]
		if curr != prev {
			strings[j] = curr
			prev = curr
			j++
		}
	}
	return strings[:j]
}
func main() {
	a := [...]string {"abc", "bcd", "bcd", "cdd", "cdd", "c", "d"}
	s := a[:]
	fmt.Printf("%s\n", s)
	s = removeAdjacentDupStringsInPlace(s)
	fmt.Printf("%s\n", s)
	b := [...]string {"\u4e16\u754c", "\u4e16\u754c", "\u4e16"}
	s = b[:]
	fmt.Printf("%s\n", s)
	s = removeAdjacentDupStringsInPlace(s)
	fmt.Printf("%s\n", s)}