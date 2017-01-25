// Examines https://github.com/golang/go/wiki/Range#gotchas
package main

import (
	"fmt"
	"math/rand"
)


func main() {
	doc := make([]map[string]int, 5) // document is a 5-paragraph slice, each of which is a map of parts of speech
	for i := 0; i < len(doc); i++ {
	// don't use range here: for _, paragraph := range(doc) {
		doc[i] = make(map[string]int)
		// now doc[i] is allocated fully, populate it as needed
		paragraph(doc[i])
	}
}
func printDoc(doc []map[string]int) {
	for _, para := range doc {
		fmt.Printf("%v\n", para)
	}
}
func paragraph(parts map[string]int) {
	parts["verb"] = rand.Intn(10)
	parts["noun"] = rand.Intn(10)
	parts["pronoun"] = rand.Intn(5)
	parts["article"] = rand.Intn(10)
}
