// Demonstrates that the dot notation is overloaded in that it works on struct as well as a pointer to it
package main

import (
	"fmt"
)

type Employee struct {
	ID int
	name string
}
var dilbert Employee
var none Employee
var boss Employee

func init() {
	dilbert.ID = 1
	dilbert.name = "Dilbert"
	boss.ID = 2
	boss.name = "Pointy-haired Boss"
}
func employeeByID1(id int) *Employee {
	if id == 1 {
		return &dilbert
	}
	if id == 2 {
		return &boss
	}
	return &none
}
func employeeByID2(id int) Employee {
	if id == 1 {
		return dilbert
	}
	if id == 2 {
		return boss
	}
	return none
}
func main() {
	fmt.Printf("%v\n", dilbert)
	var ptr *Employee = &dilbert
	ptr.name = "Dilbert Adams" // dot notation (no ->) with the ptr to struct
	fmt.Printf("%v\n", dilbert)
	// or
	fmt.Printf("%v\n", *ptr)
	//employeeByID2(2).name = "" // not ok, compile error: cannot assign to employeeByID2(2).name
	employeeByID1(2).name = "" // ok!
	fmt.Printf("%v\n", employeeByID1(2)) // boss's name wiped out

	type Empty struct {}
	var e Empty
	set := make(map[string]Empty)
	set["foo"] = e
	_, ok := set["bar"]
	if !ok {
		fmt.Println("bar is not in the set!")
	}
	_, ok = set["foo"]
	if ok {
		fmt.Println("but foo is in the set!")
	}
	type SomeKey struct {
		x int
		n string
	}
	s := make(map[SomeKey]string)
	k1 := SomeKey{x: 5, n: "foo"}
	k2 := SomeKey{x: 6, n: "foo"}
	s[k1] = k1.n
	s[k2] = k2.n
	if _, ok = s[k1]; ok {
		fmt.Printf("Key %v exists\n", k1)
	}
	k1.n = "bar"
	if _, ok = s[k1]; ok {
		fmt.Printf("Key %v still exists\n", k1)
	} else {
		fmt.Printf("OMG! mutable key %v\n", k1)
	}
}
