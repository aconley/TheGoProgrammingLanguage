package permutations

import "sort"

// This generates integer permutations lexicographically
//  using Algorithm L from Knuth vol 4, section 7.2.1.2
// This is the basic (un-optimized) version

// NextLexicographicPermutation modifies it's argument
//  to the next lexicographic permutation, returning
//  true if this was possible and false otherwise.
// Note that the elements are not required to be distinct,
//  but that interchanges of non-distinct elements are
//  not considered different.  Thus, the permuatations
//  of 113 are 113, 131, 311 only.
func NextLexicographicPermutation(a sort.Interface) bool {
	n := a.Len()
	if n == 0 {
		return false
	}

	j := n - 2
	for j >= 0 && !a.Less(j, j+1) {
		j--
	}

	// That was the final permutation
	if j < 0 {
		return false
	}

	// j is now the smallest subscript such that we've already
	//  visited all permutations beginning with a[0]..a[j].
	//  So we will now have to make a[j] larger
	// Figure out what to exchange it with.
	l := n - 1
	for !a.Less(j, l) {
		l--
	}
	a.Swap(j, l)

	// And reverse a[j+1] to a[l], which is the lexicographically
	//  least way to finish out the permutation
	k := j + 1
	l = n - 1
	for k < l {
		a.Swap(k, l)
		k++
		l--
	}

	return true
}

// Lexicographic visits all permutations of data in
//  lexicographic order, calling visit on each
//  permutation unless stopped early
func Lexicographic(data SortableInterface) {
  n := data.Len()
  if n == 0 {
    return
  }

  if !sort.IsSorted(data) {
    sort.Sort(data)
  }

  // visit the first one
  if !data.Visit() {
    // quit early
    return
  }

  for NextLexicographicPermutation(data) {
    if !data.Visit() {
      // quit early
      return
    }
  }
}