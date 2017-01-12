// http://stackoverflow.com/questions/21719769/passing-an-array-as-an-argument-in-golang

package main

import "fmt"

type name struct {
	X string
}

func main() {
	var a [3]name
	a[0] = name{"Abbed"}
	a[1] = name{"Ahmad"}
	a[2] = name{"Ghassan"}

	nameReader(a)
}

func nameReader(array []name) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i].X)
	}
}