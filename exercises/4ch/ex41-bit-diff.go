// Write a function that counts the number of bits that are different in two SHA256 hashes. (See PopCount from Section 2.6.2.)

package main

import (
	"crypto/sha256"
	"exercises/ch2/popcount"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	// number of bits that are different in c1 and c2 = number of 1 bitss in c1^c2
	nbits1 := 0
	nbits2 := 0
	nbits3 := 0
	size := len(c1)
	for i := 0; i < size; i++ {
		nbits1 += popcount.PopCount(uint64(c1[i]^c2[i]))
		nbits2 += popcount.PopCountBitwise(uint64(c1[i]^c2[i]))
		nbits3 += popcount.PopCountLoop(uint64(c1[i]^c2[i]))
	}
	fmt.Printf("Of 256, sha's differ in %d bits\n", nbits1)
	fmt.Printf("Of 256, sha's differ in %d bits\n", nbits2)
	fmt.Printf("Of 256, sha's differ in %d bits\n", nbits3)
}

