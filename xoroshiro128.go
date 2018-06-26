// Package prng implements xoroshiro128+ pseudorandom number generators.
package prng

import (
	"math/rand"

	"github.com/ericyan/prng/internal/splitmix64"
)

// A Xoroshiro128Plus is an instance of xoroshiro128+ generator.
type Xoroshiro128Plus struct {
	state [2]uint64
}

// New creates and returns a new PRNG. The seed should be two unsigned
// 64-bit integers and must not be everywhere zero.
func NewXoroshiro128Plus(seed [2]uint64) rand.Source64 {
	return &Xoroshiro128Plus{seed}
}

// Seed initializes the generator to a deterministic state. It seeds a
// SplitMix64 generator using the provided 64-bit seed value and then
// uses the outputs to initialize the xoroshiro128+ generator demanding
// 128 bits of state.
func (g *Xoroshiro128Plus) Seed(seed int64) {
	r := splitmix64.New(seed)

	g.state[0], g.state[1] = r.Uint64(), r.Uint64()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (g *Xoroshiro128Plus) Uint64() uint64 {
	s0, s1 := g.state[0], g.state[1]
	result := s0 + s1

	s1 ^= s0
	g.state[0] = rotl(s0, 55) ^ s1 ^ (s1 << 14)
	g.state[1] = rotl(s1, 36)

	return result
}

// Int63 returns a non-negative pseudorandom 63-bit integer as an int64.
func (g *Xoroshiro128Plus) Int63() int64 {
	return int63(g.Uint64())
}
