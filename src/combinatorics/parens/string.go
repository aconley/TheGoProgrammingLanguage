package parens

type stringParens struct {
	n       int
	m       int
	hasMore bool
	state   []byte
}

// CreateString makes a new nested parenthesis generator
//  based on string manipulation
func CreateString(n int) Generator {
	if n <= 0 {
		return new(stringParens)
	}
	state := make([]byte, 2*n+1, 2*n+1)
	for k := 1; k <= n; k++ {
		state[2*k-1] = '('
		state[2*k] = ')'
	}
	state[0] = ')' // Sentinel
	return &stringParens{n, 2*n - 1, true, state}
}

func (s *stringParens) HasNext() bool {
	if s == nil {
		return false
	}
	return s.hasMore
}

func (s *stringParens) GetNext() string {
	if s == nil || !s.hasMore {
		panic("Generator is done")
	}
	// Get string representation of current state
	res := string(s.state[1:])

	// Generate new state
	s.state[s.m] = ')'
	var j int
	if s.state[s.m-1] == ')' {
		s.state[s.m-1] = '('
		s.m--
	} else {
		j = s.m - 1
		for k := 2*s.n - 1; s.state[j] == '('; {
			s.state[j] = ')'
			s.state[k] = '('
			j--
			k -= 2
		}
		if j == 0 {
			s.hasMore = false
		} else {
			s.state[j] = '('
			s.m = 2*s.n - 1
		}
	}

	return res
}
