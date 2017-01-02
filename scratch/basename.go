// Implements a function similar to the basename command on Unix
package main

import "fmt"

func main() {
	fmt.Printf("%q\n", basename("foo"))
	fmt.Printf("%q\n", basename("foo.x"))
	fmt.Printf("%q\n", basename("/foo.x"))
	fmt.Printf("%q\n", basename("/a/b/foo.x"))
	fmt.Printf("%q\n", basename("/a/b/foo."))
	fmt.Printf("%q\n", basename("*.z"))
	fmt.Printf("%q\n", basename("."))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
