package main

import (
	"fmt"
	"os"
	"strconv"
)

// Rotate left in one pass.  Uses a temporary
// array
func rotateLeftOnePass(a []byte, l int) {
	if l < 0 {
		panic("Invalid (negative) left rotate amount")
	}

	n := len(a)
	if l > n {
		l = l % n
	}
	if l == 0 {
		return
	}

	tmpVal := make([]byte, l)
	copy(tmpVal, a[0:l])
	copy(a[0:], a[l:])
	copy(a[(n-l):], tmpVal)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: rotateSinglePass amount string1 ...")
	}
	rotAmt, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Couldn't get rotate amount")
	}
	fmt.Printf("Rotating by %d\n", rotAmt)
	for i := 2; i < len(os.Args); i++ {
		arg := []byte(os.Args[i])
		rotateLeftOnePass(arg, rotAmt)
		fmt.Printf("  %v\n", string(arg))
	}
}
