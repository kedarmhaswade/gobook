// Examines if all empty structs same things?
package main

import (
	"fmt"
)

type Empty struct{}
type Employee1 struct{
	ID int
	Name string
}

func main() {
	e1 := Empty{}
	e2 := Empty{}
	fmt.Printf("address of e1 = %p \n", &e1)
	fmt.Printf("address of e2 = %p \n", &e2)
	fmt.Printf("are empty structs the same object? %v \n", &e1 == &e2)
	ptr1 := &e1 // pointer to Empty
	ptr2 := &e2 // pointer to Empty
	fmt.Printf("Do both pointers to Empty struct point to the same address? %v \n", *ptr1 == *ptr2)
}