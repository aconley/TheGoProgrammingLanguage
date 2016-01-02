package popcount_test

import (
    "testing"
    "math/rand"
    
    "exercises/gopl.io/ch2/popcount"
)

const (
    randSeed int64 = 2039343
)

func getUint64() uint64 {
    x1 := uint64(rand.Uint32())
    x2 := uint64(rand.Uint32())
    return (x1 << 32) | x2
}

func BenchmarkLookup(b *testing.B) {
    rand.Seed(randSeed)
    for i := 0; i < b.N; i++ {
        popcount.PopCountByLookup(getUint64())
    }
}

func BenchmarkLookup2(b *testing.B) {
    rand.Seed(randSeed)
    for i := 0; i < b.N; i++ {
        popcount.PopCountByLookup2(getUint64())
    }
}

func BenchmarkShifting(b *testing.B) {
    rand.Seed(randSeed)
    for i := 0; i < b.N; i++ {
        popcount.PopCountByShifting(getUint64())
    }
}

func BenchmarkClearing(b *testing.B) {
    rand.Seed(randSeed)
    for i := 0; i < b.N; i++ {
        popcount.PopCountByClearing(getUint64())
    }
}

func BenchmarkHacker(b *testing.B) {
    rand.Seed(randSeed)
    for i := 0; i < b.N; i++ {
        popcount.PopCountHacker(getUint64())
    }
}
