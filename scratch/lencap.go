// Examines slices
package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July",
	8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	// note: months[0] is "", the default value of string
	fmt.Printf("%d %d \n", len(months), cap(months)) // 13 13
	Q1 := months[1:4]
	fmt.Printf("%d %d \n", len(Q1), cap(Q1)) // 3 12
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]
	fmt.Printf("%d\n", cap(summer)) // (12 + 1) - 6 = 7, i.e. we could put at most 7 elements in the slice summer
	fmt.Println(summer[:7]) // no panic but [June July August September October November December]
	//fmt.Println(summer[:8]) // panic
}
