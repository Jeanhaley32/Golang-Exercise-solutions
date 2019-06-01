package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(i int) int {
	cv := 1
	pv := 0
	rv := 0
	return func(i int) int {
		if i == 0 {
			fmt.Println(pv)
		}
		if i == 1 {
			fmt.Println(cv)
		}
		rv = cv + pv
		pv = cv
		cv = rv
		return rv
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))

	}
}
