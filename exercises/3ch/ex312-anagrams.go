// Reports whether two strings are anagrams of each other
package main

import (
	"fmt"
)

func main() {
	s1 := "केदार"
	s2 := "रदाके"
	fmt.Printf("Are %s and %s anagrams of each other? : %v\n", s1, s2, same(histo(s1), histo(s2)))
	s1 = "foo"
	s2 = "off"
	fmt.Printf("Are %s and %s anagrams of each other? : %v\n", s1, s2, same(histo(s1), histo(s2)))
	s1 = "fooo"
	s2 = "ofoo"
	fmt.Printf("Are %s and %s anagrams of each other? : %v\n", s1, s2, same(histo(s1), histo(s2)))
	s1 = "\u4e16\u754c"
	s2 = "\u754c\u4e16"
	fmt.Printf("Are %s and %s anagrams of each other? : %v\n", s1, s2, same(histo(s1), histo(s2)))
}
func same(m1 map[rune]int, m2 map[rune]int) interface{} {
	for rune, val := range m1 {
		if val != m2[rune] {
			return false
		}
	}
	return len(m1) == len(m2)
}
func histo(s string) map[rune]int {
	var m = make(map[rune]int)
	for _, r := range s {
		m[r] += 1
	}
	return m
}