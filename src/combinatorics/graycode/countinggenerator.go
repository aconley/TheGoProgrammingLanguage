package graycode

import "fmt"

// CountingGenerator is a simple counting based binary Gray code generator
type CountingGenerator struct {
	n        int
	maxstate uint32
	state    uint32
}

// CreateCounting creates a new CountingGenerator as a pointer
func CreateCounting(n int) *CountingGenerator {
	if n <= 0 {
		return nil
	}
	if n > 31 {
		panic(fmt.Sprintf("User value for n %d > max value %d", n, 31))
	}
	return &CountingGenerator{n, 1 << uint(n), 0}
}

// HasNext returns true if the generator has more values
func (c *CountingGenerator) HasNext() bool {
	return c.state < c.maxstate
}

// GetNext gets the next value
func (c *CountingGenerator) GetNext() uint32 {
	if c == nil || !c.HasNext() {
		panic("Generator is done")
	}
	res := c.state ^ (c.state >> 1)
	c.state++
	return res
}
