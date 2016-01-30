package parens

import (
	"bytes"
	"fmt"
)

const (
	mu0 uint32 = 0x55555555
	mu1 uint32 = 0x33333333
	mu2 uint32 = 0x0f0f0f0f
)

// Bitwise is a ParenGenerator using bitwise operations
type bitwise struct {
	n      int
	done   bool
	mask   uint32
	state  uint32
	buffer bytes.Buffer
}

// CreateBitwise creates a new Bitwise
func CreateBitwise(n int) Generator {
	if n > 16 {
		panic(fmt.Sprintf("Invalid size %d; must be <= 16", n))
	}
	if n <= 0 {
		return nil
	}
	nu := uint32(n)
	var b bytes.Buffer
	return &bitwise{n, false, 1 << (2*nu + 1), (1 << nu) - 1, b}
}

// HasNext returns true if there are more patterns to generate
func (b *bitwise) HasNext() bool {
	if b == nil {
		return false
	}
	return !b.done
}

// GetNext returns the next set of parenthesis, and updates
//  the state of b
func (b *bitwise) GetNext() string {
	if b == nil || b.done {
		panic("Generator is done")
	}
	// Generate new state
	res := generateFromState(b.n, b.state, &b.buffer)
	t := b.state ^ mu0
	u := t ^ (t - 1)
	v := b.state | u
	w := v + 1
	s := popCount(u & mu0)
	wp := (v & (^w)) >> s
	newstate := w + wp
	b.done = (newstate & b.mask) != 0
	b.state = newstate
	return res
}

func generateFromState(n int, v uint32, buffer *bytes.Buffer) string {
	buffer.Reset()
	for i := uint(2 * n); i > 0; i-- {
		if (1<<(i-1))&v != 0 {
			buffer.WriteByte(')')
		} else {
			buffer.WriteByte('(')
		}
	}
	return buffer.String()
}

// This is the Hacker's Delight popcount
func popCount(x uint32) uint32 {
	x = x - ((x >> 1) & mu0)
	x = (x & mu1) + ((x >> 2) & mu1)
	x = (x + (x >> 4)) & mu2
	x = x + (x >> 8)
	x = x + (x >> 16)
	return x & 0x3f
}
