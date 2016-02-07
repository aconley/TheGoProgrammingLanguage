package graycode

import "testing"

const (
	ntest = 16
)

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := CreateBasic(ntest)
		CountNSets(g)
	}
}

func BenchmarkCounting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := CreateCounting(ntest)
		CountNSets(g)
	}
}
