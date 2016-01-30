package parens

// Generator generates valid sets of nested parenthesis.
//  The interface is similar to an iterator in Java
type Generator interface {
	HasNext() bool
	GetNext() string
}
