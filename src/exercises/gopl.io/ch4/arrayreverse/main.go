package main

import "fmt"

// Reverse the array in place
// This seems to require a fixed length array
func reverse(p *[16]int) {
	for i, j := 0, 15; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}

func main() {
	var a [16]int
	for i := 0; i < 16; i++ {
		a[i] = i
	}
	fmt.Printf("Original array %v\n", a)
	reverse(&a)
	fmt.Printf("Reversed %v\n", a)
}
