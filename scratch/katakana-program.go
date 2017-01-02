package main

import "fmt"

func main() {
	// "program" in Japanese katakana
	s := "プログラム"
	// watch the space between % and x to write a space after every byte!
	fmt.Printf("%  x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s) // the rune conversion, the returned slice contains the runes, not bytes from s
	fmt.Printf("%x\n", r)  // "[30d7 30ed 30b0 30e9 30e0]"
	fmt.Printf("%s\n", string(r)) // "プログラム"
	fmt.Println(string(0x4eac)) // 京
}
