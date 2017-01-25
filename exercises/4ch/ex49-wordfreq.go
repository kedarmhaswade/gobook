// Finds the frequency of each word in the input
package main

import (
	"os"
	"bufio"
	"fmt"
)


func getWordFrequency(reader *os.File) map[string]int {
	input := bufio.NewScanner(reader)
	input.Split(bufio.ScanWords)
	wordmap := make(map[string]int)
	for input.Scan() {
		word := input.Text()
		wordmap[word] += 1
	}
	return wordmap
}
func main() {
	fmt.Println("Word histogram")
	for w, f := range getWordFrequency(os.Stdin) {
		fmt.Printf("%s\t%d\n", w, f)
	}
}
