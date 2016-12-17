// Provides a rudimentary front end for the unit conversion program
package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"exercises/ch2/convert"
)

func main() {
	var fromUnit string // lower case name of the unit to convert from
	var toUnit string // lower case name of the unit to convert to
	var quantity float64 = 1.0 // the quantity to convert, 1.0 by default

	nargs := len(os.Args) - 1 // program name is stripped
	println(nargs)
	if nargs == 3 {
		fromUnit = os.Args[1]
		toUnit = os.Args[2]
		quantity, _ = strconv.ParseFloat(os.Args[3], 64)
		fromToQty(fromUnit, toUnit, quantity)
	} else if nargs == 3 {
		fromUnit = os.Args[1]
		toUnit = os.Args[2]
		fromToQty(fromUnit, toUnit, 1)
	} else if nargs == 0 {
		fromToQty("in", "m", 1)
		fromToQty("m", "in", 1)
	}
}
// Converts the given quantity from the from unit to the to unit
func fromToQty(from string, to string, qty float64) {

	if (strings.HasPrefix(from, "in") || strings.HasPrefix(from, "inch")) &&
		(to == "m" || strings.HasPrefix(to, "mtr") || strings.HasPrefix(to, "meter") || strings.HasPrefix(to, "metre")) {
		fmt.Printf("%.3f %-5s = %.3f %5s\n", qty, from, convert.InchToMeter(qty), to)
	}
	if (strings.HasPrefix(to, "in") || strings.HasPrefix(to, "inch")) &&
		(from == "m" || strings.HasPrefix(from, "mtr") || strings.HasPrefix(from, "meter") || strings.HasPrefix(from, "metre")) {
		fmt.Printf("%.3f %-5s = %.3f %5s\n", qty, from, convert.MeterToInch(qty), to)
	}
}

