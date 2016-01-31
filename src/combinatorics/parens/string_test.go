package parens

import (
	"fmt"
	"testing"
)

func TestDytinhSmall(t *testing.T) {
	t.Log("Testing small string example (2 parens)")
	g := CreateString(2)
	if !g.HasNext() {
		t.Errorf("2 paren set should not be empty on initialization")
	}
	if s := g.GetNext(); s != "()()" {
		t.Errorf("First one should be (()), was %s", s)
	}
	if !g.HasNext() {
		t.Errorf("2 paren set should have second set")
	}
	if s := g.GetNext(); s != "(())" {
		t.Errorf("Second one should be ()(), was %s", s)
	}
}

func TestMedium(t *testing.T) {
	t.Log("Testing medium example (4 parens)")

	var expected = [14]string{
		"()()()()", "()()(())", "()(())()", "()(()())",
		"()((()))", "(())()()", "(())(())", "(()())()",
		"(()()())", "(()(()))", "((()))()", "((())())",
		"((()()))", "(((())))"}
	n := len(expected)
	g := CreateString(4)
	var cntr int
	for cntr = 0; g.HasNext(); cntr++ {
		if cntr >= n {
			t.Errorf("Got more nested parens than expected (%d)", n)
		}
		s := g.GetNext()
		if s != expected[cntr] {
			t.Errorf("Expected %s for %dth set, got %s",
				expected[cntr], cntr+1, s)
		}
	}
	if cntr != n {
		t.Errorf("Expected %d sets of parens, got %d", n, cntr)
	}
}

func countNSets(g Generator) int {
	var cntr int
	for cntr = 0; g.HasNext(); cntr++ {
		_ = g.GetNext()
	}
	return cntr
}

func TestStringCount(t *testing.T) {
	// Here we just test that we get the expected number
	t.Log("Doing number of entries test for bitwise")
	n := [...]int{5, 6, 7, 8}
	nexpected := [...]int{42, 132, 429, 1430}
	for i, v := range n {
		t.Log(fmt.Sprintf("Running test on size %d", v))
		g := CreateBitwise(v)
		nFound := countNSets(g)
		if nFound != nexpected[i] {
			t.Errorf("Expected %d sets for %d, got %d",
				nexpected[i], v, nFound)
		}
	}
}
