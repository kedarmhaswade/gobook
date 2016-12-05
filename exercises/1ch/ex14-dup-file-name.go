// Exercise 1.4: Modify dup2 to print the names of all files
// in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		counts["Stdin"] = make(map[string]int)
		countLines(os.Stdin, counts["Stdin"])
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex14: %v\n", err)
				continue
			}
			counts[arg] = make(map[string]int)
			countLines(f, counts[arg])
			f.Close()
		}
	}
	for filename, m := range counts {
		fmt.Println(filename)
		for line, count := range m {
			if count > 1 {
				fmt.Printf("%d\t%s\n", count, line)
			}
		}
	}
}
func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}
