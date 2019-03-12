package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64, count int) {
	
	z = 1.0
	prev := z + 1e-7
	count = 1

	for i := &count ; math.Abs(z-prev) > 1e-8; *i++ {
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	return
}

func main() {
	const target = 7

	z, count := Sqrt(target)
	fmt.Println("for", count, "tries, we found the "+
		"sqrt is", z)

	fmt.Println("the real sqrt is", math.Sqrt(target))

}
