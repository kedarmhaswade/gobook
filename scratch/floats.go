// Experiments with
package main

import (
	"fmt"
	"math"
	//"strconv"
	"strconv"
)

func main() {
	var f float32 = 1 << 24
	fmt.Printf("%032b\n", math.Float32bits(f))
	f++
	fmt.Printf("%032b\n", math.Float32bits(f))
	i, _ := strconv.ParseInt("10010111", 2, 32)
	fmt.Printf("%d\n", i)


	for f = 1.0; f < 10; f++ {
		fmt.Printf("%f %032b\n", f, math.Float32bits(f))
	}
	fmt.Printf("%032b\n", math.Float32bits(1.0/3))
	//f = (1 << 24) - 1
	//oldf := f
	//for true {
	//	if f != (f + 1) {
	//		fmt.Printf("Got the precision right for: %f\n", f)
	//		break
	//	}
	//	f++
	//	fmt.Printf("%032b\n", math.Float32bits(f))
	//}
	//fmt.Printf("after %d increments!", int(f - oldf))
}
