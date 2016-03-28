package permutations

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
func NextLexicographicPermutation(a []int) bool {
	n := len(a)
	if n == 0 {
		return false
	}

	j := n - 2
	for j >= 0 && a[j] >= a[j+1] {
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
	aj := a[j]
	for aj >= a[l] {
		l--
	}
	a[j], a[l] = a[l], a[j]

	// And reverse a[j+1] to a[l], which is the lexicographically
	//  least way to finish out the permutation
	k := j + 1
	l = n - 1
	for k < l {
		a[k], a[l] = a[l], a[k]
		k++
		l--
	}

	return true
}
