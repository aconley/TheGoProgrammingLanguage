package graycode

import "testing"

func TestCountingGray3(t *testing.T) {
	t.Log("Testing small (3) Gray Binary Code")
	b := CreateCounting(3)
	var expected = [8]uint32{0, 1, 3, 2, 6, 7, 5, 4}
	n := len(expected)
	var cntr int
	for cntr = 0; b.HasNext(); cntr++ {
		if cntr >= n {
			t.Errorf("Got more values than expected (%d)", n)
		}
		s := b.GetNext()
		if s != expected[cntr] {
			t.Errorf("Expected %d for %dth value, got %d",
				expected[cntr], cntr+1, s)
		}
	}
	if cntr != n {
		t.Errorf("Expected %d values , got %d", n, cntr)
	}
}

func TestCountingGray4(t *testing.T) {
	t.Log("Testing small (4) Gray Binary Code")
	b := CreateCounting(4)
	var expected = [16]uint32{0, 1, 3, 2, 6, 7, 5, 4, 12,
		13, 15, 14, 10, 11, 9, 8}
	n := len(expected)
	var cntr int
	for cntr = 0; b.HasNext(); cntr++ {
		if cntr >= n {
			t.Errorf("Got more values than expected (%d)", n)
		}
		s := b.GetNext()
		if s != expected[cntr] {
			t.Errorf("Expected %d for %dth value, got %d",
				expected[cntr], cntr+1, s)
		}
	}
	if cntr != n {
		t.Errorf("Expected %d values , got %d", n, cntr)
	}
}
