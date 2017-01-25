// Examines the anonymous fields of structs
package main

import "fmt"

type Point struct {
	X float64
	Y float64
}
type Line struct {
	Point // first point, say the default point
	P2 Point // second point
}
func main() {

	p1 := Point{1, 2}
	p2 := Point{3, 4}
	line := Line{Point: p1, P2: p2}
	fmt.Printf("%f\n", line.X)
	fmt.Printf("%f\n", line.Y)
	// Point, i.e. the first point defaults to origin for this line, line1
	// since that value is the zero value for the field "Point"
	line1 := Line{P2: p2}
	fmt.Printf("%f\n", line1.X)
	fmt.Printf("%f\n", line1.Y)
}

func len(line Line)