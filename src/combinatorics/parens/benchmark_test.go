package parens

import "testing"

const (
	ntest = 8
)

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := CreateString(ntest)
		CountNSets(g)
	}
}

func BenchmarkBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := CreateBitwise(ntest)
		CountNSets(g)
	}
}
