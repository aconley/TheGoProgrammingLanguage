package parens

// ParenGenerator generates valid sets of nested parenthesis
type ParenGenerator interface {
	HasNext() bool
	GetNext() string
}
