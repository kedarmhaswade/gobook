// Given a list (slice) of strings, return the nonempty ones
package main

import "fmt"

func nonemptyNotInPlace(strings []string) []string {
	var ne []string
	for _, str := range strings {
		if len(str) > 0 {
			ne = append(ne, str)
		}
	}
	return ne
}
// In-place update of the given slice's underlying array
func nonEmptyInplace(strings []string) []string {
	i := 0 // used as the index to mutate the underlying array of passed slice
	for _, str := range strings { // the implicit index is ignored!
		if str != "" {
			strings[i] = str
			i++
		}
	}
	return strings[:i]
}
func nonEmptyUsingAppend(strings []string) []string {
	z := strings[:0]
	for _, s := range strings {
		if s != "" {
			z = append(z, s)
		}
	}
	return z
}
func main() {
	strs := [...]string {"", "foo", "bar"}
	fmt.Printf("%q\n", nonEmptyInplace(strs[:]))
	fmt.Printf("%q\n", strs)
}
