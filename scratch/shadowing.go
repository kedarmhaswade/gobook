// Shadowing of a package-level variable by the init function
package main

import (
	"os"
	"log"
)

var cwd string

func init() {
	var lcwd string
	var err error
	lcwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("error: %v, exiting\n", err)
	}
	cwd = lcwd
}

func main() {
	log.Printf("cwd is %v\n", cwd)
}
