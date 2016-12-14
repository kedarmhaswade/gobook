package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5/9)
}
func main() {
	var c Celsius
	var f Fahrenheit
	fmt.Printf("%v\n", (0 == 0.0)) // true
	// fmt.Printf("%v\n", (c == f)) // mismatched types
	fmt.Printf("%v\n", c == Celsius(f)) // type conversion, value is left intact!
}

