package main

import "fmt"

func main() {
	var x int32
	print(int(x))
}

func print(x int) {
	fmt.Printf("%d\n", x)
}
