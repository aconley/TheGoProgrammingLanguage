package permutations

import (
	"sort"
	"testing"
)

// Tests of NextLexicographicPermutation

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

// Tests of the visitor version

// Counting visitor
type countingIntVisitor struct {
	s       []int
	n       int
}

func (s *countingIntVisitor) Len() int {
    return len(s.s)
}
func (s *countingIntVisitor) Swap(i, j int) {
    s.s[i], s.s[j] = s.s[j], s.s[i]
}
func (s *countingIntVisitor) Less(i, j int) bool {
    return s.s[i] < s.s[j]
}

func (s *countingIntVisitor) Visit() bool {
  s.n++
  return true
}

func TestLexicographicVisitor4count(t *testing.T) {
	t.Log("Counting visitor based permutations of 1223")
	testVal := &countingIntVisitor{s: []int{1, 2, 3, 4}, n: 0}
  Lexicographic(testVal);
  if testVal.n != 24 {
    t.Errorf("Expected %d perms, got %d", 24, testVal.n)
  }
}

// Record visits
type recordIntVisitor struct {
	s       []int
	visited [][]int
}

func (s *recordIntVisitor) Len() int {
    return len(s.s)
}
func (s *recordIntVisitor) Swap(i, j int) {
    s.s[i], s.s[j] = s.s[j], s.s[i]
}
func (s *recordIntVisitor) Less(i, j int) bool {
    return s.s[i] < s.s[j]
}

func (s *recordIntVisitor) Visit() bool {
  newval := make([]int, len(s.s))
  copy(newval, s.s)
  s.visited = append(s.visited, newval)
  return true
}

func (s *recordIntVisitor) NVisited() int {
  return len(s.visited)
}

func TestLexicographicVisitor4rep(t *testing.T) {
	t.Log("Testing visitor based permutations of 1223")
	var expected = [12][4]int{{1, 2, 2, 3}, {1, 2, 3, 2}, {1, 3, 2, 2},
		{2, 1, 2, 3}, {2, 1, 3, 2}, {2, 2, 1, 3}, {2, 2, 3, 1},
		{2, 3, 1, 2}, {2, 3, 2, 1}, {3, 1, 2, 2}, {3, 2, 1, 2},
		{3, 2, 2, 1}}
	n := len(expected)

	testVal := &recordIntVisitor{s: []int{1, 2, 2, 3}, visited: nil}

  Lexicographic(testVal);

  if testVal.NVisited() != n {
    t.Errorf("Expected %d perms, got %d", n, testVal.NVisited())
  }

  for i := 0; i < n; i++ {
    if !compareIntSlice(expected[i][:], testVal.visited[i]) {
      t.Errorf("On permutation %d expected %v got %v", i,
        expected[i], testVal.visited[i])
    }
  }
}

func compareIntSlice(v1 []int, v2 []int) bool {
  if len(v1) != len(v2) {
    return false
  }
  for i := range(v1) {
    if v1[i] != v2[i] {
      return false
    }
  }
  return true
}

// Benchmark

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
