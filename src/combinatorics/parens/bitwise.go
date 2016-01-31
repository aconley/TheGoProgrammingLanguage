package parens

import (
	"fmt"
)

const (
	mu0 uint32 = 0x55555555
	mu1 uint32 = 0x33333333
	mu2 uint32 = 0x0f0f0f0f
)

// Bitwise is a ParenGenerator using bitwise operations
type bitwise struct {
	n       int
	hasMore bool
	mask    uint32
	state   uint32
	buffer  []byte
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
	b := make([]byte, 2*n, 2*n)
	return &bitwise{n, true, 1 << (2*nu + 1), (1 << nu) - 1, b}
}

// HasNext returns true if there are more patterns to generate
func (b *bitwise) HasNext() bool {
	if b == nil {
		return false
	}
	return b.hasMore
}

// GetNext returns the next set of parenthesis, and updates
//  the state of b
func (b *bitwise) GetNext() string {
	if b == nil || !b.hasMore {
		panic("Generator is done")
	}
	// Get the string representation of the current state
	res := generateFromState(b.n, b.state, b.buffer)

	// Generate new state
	t := b.state ^ mu0
	u := t ^ (t - 1)
	v := b.state | u
	w := v + 1
	s := popCount(u & mu0)
	wp := (v & (^w)) >> s
	newstate := w + wp

	b.hasMore = (newstate & b.mask) == 0
	b.state = newstate
	return res
}

// generatesFromState converts the internal state to a string
//  representation
func generateFromState(n int, v uint32, buffer []byte) string {
	var ival uint32 = 1 << uint(2*n-1)
	for i := 0; i < 2*n; i++ {
		if (ival>>uint(i))&v == 0 {
			buffer[i] = '('
		} else {
			buffer[i] = ')'
		}
	}
	return string(buffer)
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
