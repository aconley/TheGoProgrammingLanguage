package permutations

import (
	"sort"
	"testing"
)

func TestBasicLexicographic3(t *testing.T) {
	t.Log("Testing permutations of 123")
	var expected = [6][3]int{{1, 2, 3}, {1, 3, 2},
		{2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	n := len(expected)
	var testVal = sort.IntSlice(expected[0][:])
	var ok bool
	for cntr := 1; cntr < n; cntr++ {
		ok = NextLexicographicPermutation(testVal)
		if !ok {
			t.Errorf("Expected to be able to get next perm at pos %d",
				cntr)
		}
		for i, v := range expected[cntr] {
			if v != testVal[i] {
				t.Errorf("In perm %d at position %d expected %d got %d",
					cntr, i, v, testVal[i])
			}
		}
	}
	ok = NextLexicographicPermutation(testVal)
	if ok {
		t.Errorf("Shouldn't be able to get any more permutations")
	}
}

func TestBasicLexicographic3rep(t *testing.T) {
	t.Log("Testing permutations of 113")
	var expected = [3][3]int{{1, 1, 3}, {1, 3, 1}, {3, 1, 1}}
	n := len(expected)
	var testVal = sort.IntSlice(expected[0][:])
	var ok bool
	for cntr := 1; cntr < n; cntr++ {
		ok = NextLexicographicPermutation(testVal)
		if !ok {
			t.Errorf("Expected to be able to get next perm at pos %d",
				cntr)
		}
		for i, v := range expected[cntr] {
			if v != testVal[i] {
				t.Errorf("In perm %d at position %d expected %d got %d",
					cntr, i, v, testVal[i])
			}
		}
	}
	ok = NextLexicographicPermutation(testVal)
	if ok {
		t.Errorf("Shouldn't be able to get any more permutations")
	}
}

func TestBasicLexicographic4rep(t *testing.T) {
	t.Log("Testing permutations of 1223")
	var expected = [12][4]int{{1, 2, 2, 3}, {1, 2, 3, 2}, {1, 3, 2, 2},
		{2, 1, 2, 3}, {2, 1, 3, 2}, {2, 2, 1, 3}, {2, 2, 3, 1},
		{2, 3, 1, 2}, {2, 3, 2, 1}, {3, 1, 2, 2}, {3, 2, 1, 2},
		{3, 2, 2, 1}}
	n := len(expected)
	var testVal = sort.IntSlice(expected[0][:])
	var ok bool
	for cntr := 1; cntr < n; cntr++ {
		ok = NextLexicographicPermutation(testVal)
		if !ok {
			t.Errorf("Expected to be able to get next perm at pos %d",
				cntr)
		}
		ok = true
		for i, v := range expected[cntr] {
			if v != testVal[i] {
				ok = false
			}
		}
		if !ok {
			t.Errorf("For perm %d expected %v got %v",
				cntr, expected[cntr], testVal)
		}
	}
	ok = NextLexicographicPermutation(testVal)
	if ok {
		t.Errorf("Shouldn't be able to get any more permutations")
	}
}
