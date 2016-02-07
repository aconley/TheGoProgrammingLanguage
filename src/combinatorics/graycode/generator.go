package graycode

// Generator generates gray codes.
//  The interface is similar to an iterator in Java
type Generator interface {
	HasNext() bool
	GetNext() uint32
}

// CountNSets counts the number of binary values remaining
//  in g destructively
func CountNSets(g Generator) int {
	var cntr int
	for cntr = 0; g.HasNext(); cntr++ {
		_ = g.GetNext()
	}
	return cntr
}
