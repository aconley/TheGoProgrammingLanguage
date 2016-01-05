// wodfreq computes counts of strings
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	maxStrLen = 15
)

func main() {
	counts := make(map[string]int) // counts of Unicode strings
	var wordlen [maxStrLen + 1]int // count of lengths of strings
	longer := 0                    // count of strings longer than maxStrLen

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		word := strings.ToLower(in.Text())
		counts[word]++
		n := len(word)
		if n > maxStrLen {
			longer++
		} else {
			wordlen[n]++
		}
	}

	fmt.Printf("word\tcount\n")
	for s, n := range counts {
		fmt.Printf("%s\t%d\n", s, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range wordlen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if longer > 0 {
		fmt.Printf("\n%d strings longer than %d\n", longer, maxStrLen)
	}
}
