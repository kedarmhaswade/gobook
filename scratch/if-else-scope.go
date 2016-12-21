// Demonstrates the use of declaration in if and its scope
package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	if x := rand.Float32(); x < 1.0/3 { // x is declared here! Note the semicolon!
		fmt.Printf("x: %v is in the first one third\n", x)
	} else if x < 2.0/3 {
		fmt.Printf("x: %v is in the second one third\n", x) // the x declared in if clause is in scope here.
	} else {
		fmt.Printf("x: %v is in the last one third\n", x) // the x declared in if clause is in scope here.
	}
	//fmt.Println("Once again, x = %v\n", x) // compiler error: undefined: x
}


