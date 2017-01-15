package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3
	fmt.Printf("%d, %d, %d\n", a, b, c)
	var runes []rune
	for _, r := range "Hello, \u4e16\u754c" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']

	x := false
	fmt.Printf("%v %v\n", &x, *&x)
	var y *bool  = new(bool)
	fmt.Printf("%v %[1]v\n", y)
}

