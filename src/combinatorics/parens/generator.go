package parens

// Generator generates valid sets of nested parenthesis.
//  The interface is similar to an iterator in Java
type Generator interface {
	HasNext() bool
	GetNext() string
}

// CountNSets counts the number of parens remaining
//  in g destructively
func CountNSets(g Generator) int {
	var cntr int
	for cntr = 0; g.HasNext(); cntr++ {
		_ = g.GetNext()
	}
	return cntr
}
