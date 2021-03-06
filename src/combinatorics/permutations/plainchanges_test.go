package permutations

import (
  "testing"
)

// Tests of PlainChanges

func TestPlainChangesVisitor4rep(t *testing.T) {
	t.Log("Testing visitor based plain permutations of 1234")
	var expected = [24][4]int{
    {1, 2, 3, 4}, {1, 2, 4, 3}, {1, 4, 2, 3}, {4, 1, 2, 3},
    {4, 1, 3, 2}, {1, 4, 3, 2}, {1, 3, 4, 2}, {1, 3, 2, 4},
    {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 4, 1, 2}, {4, 3, 1, 2},
    {4, 3, 2, 1}, {3, 4, 2, 1}, {3, 2, 4, 1}, {3, 2, 1, 4},
    {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 3, 1}, {4, 2, 3, 1},
    {4, 2, 1, 3}, {2, 4, 1, 3}, {2, 1, 4, 3}, {2, 1, 3, 4}}
	n := len(expected)

	testVal := &recordIntVisitor{s: []int{1, 2, 3, 4}, visited: nil}

  PlainChanges(testVal);

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

func TestPlainChanges8Count(t *testing.T) {
  t.Log("Testing number of plain change permutations of 8 entries")
  testVal := &countingIntVisitor{s: []int{0, 1, 2, 3, 4, 5, 6, 7}, n: 0}
  PlainChanges(testVal)

  if testVal.n != 40320 {
    t.Errorf("Got unexpected number of plain permutations: %d expected %d",
      testVal.n, 40320)
  }
}

// Benchmark the number of permutations of
//  [0, 9] using the countering permuter
func BenchmarkPlain9(b *testing.B) {
  a0 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
  a := make([]int, len(a0))
  for n := 0; n < b.N; n++ {
    copy(a, a0)
    testVal := &countingIntVisitor{s: a, n: 0}
    PlainChanges(testVal)
  }
}