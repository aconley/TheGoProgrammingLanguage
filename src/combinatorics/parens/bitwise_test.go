package parens

import (
	"fmt"
	"testing"
)

func TestBitwiseSmall(t *testing.T) {
	t.Log("Testing small bitwise example (2 parens)")
	g := CreateBitwise(2)
	if !g.HasNext() {
		t.Errorf("2 paren set should not be empty on initialization")
	}
	if s := g.GetNext(); s != "(())" {
		t.Errorf("First one should be (()), was %s", s)
	}
	if !g.HasNext() {
		t.Errorf("2 paren set should have second set")
	}
	if s := g.GetNext(); s != "()()" {
		t.Errorf("Second one should be ()(), was %s", s)
	}
}

func TestBitwiseMedium(t *testing.T) {
	t.Log("Testing medium bitwise example (4 parens)")

	var expected = [14]string{
		"(((())))", "((()()))", "((())())", "((()))()",
		"(()(()))", "(()()())", "(()())()", "(())(())",
		"(())()()", "()((()))", "()(()())", "()(())()",
		"()()(())", "()()()()"}
	n := len(expected)
	g := CreateBitwise(4)
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

func TestBitwiseCount(t *testing.T) {
	t.Log("Doing number of entries test for bitwise")
	// Here we just test that we get the expected number
	n := [...]int{5, 6, 7, 8}
	nexpected := [...]int{42, 132, 429, 1430}
	for i, v := range n {
		t.Log(fmt.Sprintf("\tRunning test on size %d", v))
		g := CreateBitwise(v)
		nFound := CountNSets(g)
		if nFound != nexpected[i] {
			t.Errorf("Expected %d sets for %d, got %d",
				nexpected[i], v, nFound)
		}
	}
}
