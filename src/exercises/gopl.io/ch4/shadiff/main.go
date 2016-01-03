package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"exercises/gopl.io/ch2/popcount"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: shadiff arg1 arg2")
		return
	}

	sha1 := sha256.Sum256([]byte(os.Args[1]))
	sha2 := sha256.Sum256([]byte(os.Args[2]))
	var ndiff int
	for i := 0; i < len(sha1); i++ {
		ndiff += popcount.PopCountHacker(uint64(sha1[i] ^ sha2[i]))
	}

	fmt.Printf("Number of different bits: %d\n", ndiff)
}
