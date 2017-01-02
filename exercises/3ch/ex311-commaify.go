// Inserts commas into the string that represents integer or float
package main

import (
	"fmt"
	"exercises/ch3/format"
)

func main() {
	fmt.Printf("%s\n", format.CommaFloat("+23456.78", true))
	fmt.Printf("%s\n", format.CommaFloat("-23456", true))
	fmt.Printf("%s\n", format.CommaFloat("1234.56", true))
	fmt.Printf("%s\n", format.CommaFloat("1234567890.12", true))
}
