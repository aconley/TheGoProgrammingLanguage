package main

import (
	"fmt"
)

func min(val1 int, vals ...int) int {
	currMin := val1
	for _, val := range vals {
		if val < currMin {
			currMin = val
		}
	}
	return currMin
}

func main() {
	fmt.Printf("The minimum of 0 is %d\n", min(0))
	fmt.Printf("The minimum of [0 1 2] is %d\n", min(0, 1, 2))
	fmt.Printf("The minimum of [1 0 2] is %d\n", min(1, 0, 2))
}
