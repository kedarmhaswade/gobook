// Uses the Fields function
package main

import (
	"strings"
	"fmt"
)

func main() {
	for _, n := range strings.Fields("larry\tmoe\tcurly") {
		fmt.Printf("%s\n", n)
	}
}
