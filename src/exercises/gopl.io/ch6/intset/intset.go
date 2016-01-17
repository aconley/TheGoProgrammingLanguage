// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"

	"exercises/gopl.io/ch2/popcount"
)

const (
	mod64 uint = 64 - 1
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x)&mod64
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x)&mod64
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds a list of values to be added
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// Remove removes the non-negative value x from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x)&mod64
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

// Clear clears all values (but does not deallocate memory)
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// DifferenceWith sets s to the differences of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// SymmetricDifferenceWith sets s to the symmetric differences of s and t
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Copy returns a copy of the IntSet
func (s *IntSet) Copy() *IntSet {
	if s == nil {
		return nil
	}
	words := make([]uint64, len(s.words))
	for i, word := range s.words {
		words[i] = word
	}
	return &IntSet{words}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements in the set
func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		len += popcount.PopCountHacker(word)
	}
	return len
}

// Elems retuns a slice of all the elements in s
func (s *IntSet) Elems() []int {
	nelems := s.Len()
	if nelems == 0 {
		return nil
	}
	retval := make([]int, nelems)
	cntr := 0
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				retval[cntr] = 64*i + j
				cntr++
			}
		}
	}
	return retval
}
