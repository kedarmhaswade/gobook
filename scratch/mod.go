package main

import "fmt"

func main() {
	// / on integer values truncates the result toward 0
	// % on integer values returns a value same whose sign is that of the dividend
	fmt.Printf("5/-3 = %d and 5%%-3 = %d\n-5/3 = %d and -5/-3 = %d\n-5%%3 = %d and -5%%-3 = %d\n", 5 / -3, 5 % -3, -5 / 3, -5 / -3, -5 % 3, -5 % -3)
}
