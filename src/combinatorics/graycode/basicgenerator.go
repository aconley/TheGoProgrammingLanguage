package graycode

import (
	"fmt"
)

// BasicGenerator is a simple Gray binary code generator
//  corresponding to Knuth 7.2.1.1 Algorithm G
type BasicGenerator struct {
	n       uint8
	hasMore bool
	a       uint32
	ainf    bool
}

// CreateBasic creates a new instance of a BasicGenerator
//  as a pointer
func CreateBasic(n int) *BasicGenerator {
	if n <= 0 {
		return nil
	}
	if n > 32 {
		panic(fmt.Sprintf("User value for n %d > max value %d", n, 32))
	}
	return &BasicGenerator{uint8(n), true, 0, false}
}

// HasNext returns true if the generator has more values
func (b *BasicGenerator) HasNext() bool {
	if b == nil {
		return false
	}
	return b.hasMore
}

// GetNext gets the next value
func (b *BasicGenerator) GetNext() uint32 {
	if b == nil || !b.hasMore {
		panic("Generator is done")
	}
	res := b.a
	b.ainf = !b.ainf
	var j uint8

	if b.ainf {
		j = 0
	} else {
		j = ruler(b.a) + 1
	}
	if j >= b.n {
		b.hasMore = false
	} else {
		b.a ^= 1 << j
	}

	return res
}

// ruler is the binary ruler function (count number of trailing 0s)
//  from Hacker's Delight with a hack for 0 trailing 0s since
//  we are counting up and so half are odd
func ruler(x uint32) uint8 {
	if x == 0 {
		return 32
	}
	if x&1 != 0 {
		return 0
	}
	var n uint8 = 1
	if (x & 0x0000ffff) == 0 {
		n += 16
		x >>= 16
	}
	if (x & 0x000000ff) == 0 {
		n += 8
		x >>= 8
	}
	if (x & 0x0000000f) == 0 {
		n += 4
		x >>= 4
	}
	if (x & 0x00000003) == 0 {
		n += 2
		x >>= 2
	}
	return n - uint8(x&1)
}
