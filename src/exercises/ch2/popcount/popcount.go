// Various funnctions for population count of an unsigned integer
package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i & 1)
	}
}

// Uses the pc array and an unrolled loop to find the pop count of the given argument
func PopCount(n uint64) int {
	return  int(pc[byte(n >> (0 * 8))] +
		pc[byte(n >> (1 * 8))] +
		pc[byte(n >> (2 * 8))] +
		pc[byte(n >> (3 * 8))] +
		pc[byte(n >> (4 * 8))] +
		pc[byte(n >> (5 * 8))] +
		pc[byte(n >> (6 * 8))] +
		pc[byte(n >> (7 * 8))])
}

// Uses the pc array and a loop to find the pop count of the given argument
func PopCountLoop(n uint64) int {
	c := 0
	var s uint
	for s = 0; s < 8; s++ {
		c += int(pc[byte(n >> (s * 8))])
	}
	return c
}

// Uses the pc array and bitwise operator to find the pop count of the given argument.
// x & (x - 1) clears the rightmost 1 bit in x
func PopCountBitwise(n uint64) int {
	c := 0
	for n != 0 {
		n &= n - 1
		c ++
	}
	return c
}