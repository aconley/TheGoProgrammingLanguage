// Anagram checker
package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: anagram string1 string2")
	}
	s1, s2 := os.Args[1], os.Args[2]
	if areAnagrams(s1, s2) {
		fmt.Printf("%s and %s are anagrams\n", s1, s2)
	} else {
		fmt.Printf("%s and %s are not anagrams\n", s1, s2)
	}
}

func areAnagrams(s1, s2 string) bool {
	m1 := countChars(s1)
	m2 := countChars(s2)
	return reflect.DeepEqual(m1, m2)
}

func countChars(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	return m
}
