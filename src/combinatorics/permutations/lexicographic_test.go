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

func TestLexicographic8(t *testing.T) {
  t.Log("Testing number of permutations of 8 entries")
  a := sort.IntSlice([]int{0, 1, 2, 3, 4, 5, 6, 7})
  cnt := 1
  for NextLexicographicPermutation(a) {
    cnt++
  }
  if (cnt != 40320) {
    t.Errorf("Got unexpected number of permutations: %d expected %d",
      cnt, 40320)
  }
}

func BenchmarkLexicographic8(b *testing.B) {
  a0 := []int{0, 1, 2, 3, 4, 5, 6, 7}
  a := make([]int, len(a0))
  for n := 0; n < b.N; n++ {
    copy(a, a0)
    a := sort.IntSlice(a)
    cnt := 1
    for NextLexicographicPermutation(a) {
      cnt++
    }
  }
}
