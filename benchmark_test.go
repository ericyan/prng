package prng

import (
	"math/rand"
	"testing"
)

func benchmarkRandSource(s rand.Source, b *testing.B) {
	r := rand.New(s)
	r.Seed(0x0ddc0ffeebadf00d)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Uint64()
	}
	b.SetBytes(8)
}

func BenchmarkXoroshiro128Plus(b *testing.B) {
	benchmarkRandSource(new(Xoroshiro128Plus), b)
}

func BenchmarkMathRand(b *testing.B) {
	benchmarkRandSource(rand.NewSource(1), b)
}
