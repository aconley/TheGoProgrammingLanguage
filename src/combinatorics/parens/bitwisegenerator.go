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

// BitwiseGenerator is a ParenGenerator using bitwise operations
type BitwiseGenerator struct {
	n      int
	done   bool
	mask   uint32
	state  uint32
	buffer bytes.Buffer
}

// CreateBitwiseGenerator creates a new BitwiseGenerator
func CreateBitwiseGenerator(n int) BitwiseGenerator {
	if n > 16 {
		panic(fmt.Sprintf("Invalid size %d; must be <= 16", n))
	}
	nu := uint32(n)
	var b bytes.Buffer
	return BitwiseGenerator{n, false, 1 << (nu + 1), (1 << nu) - 1, b}
}

// HasNext returns true if there are more patterns to generate
func (b *BitwiseGenerator) HasNext() bool {
	return !b.done
}

// GetNext returns the next set of parenthesis, and updates
//  the state of b
func (b *BitwiseGenerator) GetNext() string {
	if b.done {
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
	for i := uint(n); n > 0; i++ {
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
