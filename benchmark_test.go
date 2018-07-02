package prng_test

import (
	"math/rand"
	"testing"

	"github.com/ericyan/prng"
	"github.com/ericyan/prng/internal/splitmix64"
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
	benchmarkRandSource(prng.NewXoroshiro128Plus(1), b)
}

func BenchmarkXoroshiro128StarStar(b *testing.B) {
	benchmarkRandSource(prng.NewXoroshiro128StarStar(1), b)
}

func BenchmarkSplitMix64(b *testing.B) {
	benchmarkRandSource(splitmix64.New(1), b)
}

func BenchmarkMathRand(b *testing.B) {
	benchmarkRandSource(rand.NewSource(1), b)
}
