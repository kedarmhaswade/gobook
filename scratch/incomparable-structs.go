// Comparable structs can be keys of maps, what about incomparable ones?
// (If all the fields of a struct are comparable, the struct itself is comparable)
package main

import "fmt"

func main() {
	type Person struct {
		name string
		phones [2]string // arrays are comparable
		//phones []string // if  phones becomes a slice, struct becomes incomparable!
	}
	type Address string
	var ab = make(map[Person]Address) // some address-book
	p1 := Person{"foo", [...]string{"123", "456"}}
	// p1 := Person{"foo", [...]string{"123", "456"}}
	p2 := Person{"bar", [...]string{"123", "456"}}
	// p2 := Person{"bar", [...]string{"123", "456"}}
	a1 := Address("10 Downing Street")
	a2 := Address("1 Havaa Mahal")
	ab[p1] = a1
	ab[p2] = a2
	//p3 := Person{"foo", []string{"123", "456"}}
	if _, ok := ab[p1]; ok {
		fmt.Printf("Yes, key %v exists\n", p1)
	}
}
