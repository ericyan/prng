// Package prng implements xoroshiro128+ pseudorandom number generators.
package prng

import "math/rand"

// A PRNG is an instance of xoroshiro128+ generator.
type PRNG struct {
	state [2]uint64
}

// New creates and returns a new PRNG. The seed should be two unsigned
// 64-bit integers and must not be everywhere zero.
func New(seed [2]uint64) rand.Source64 {
	return &PRNG{seed}
}

// Seed initializes the generator to a deterministic state. It seeds a
// SplitMix64 generator using the provided 64-bit seed value and then
// uses the outputs to initialize the xoroshiro128+ generator demanding
// 128 bits of state.
func (g *PRNG) Seed(seed int64) {
	for i := 0; i < 2; i++ {
		z := uint64(seed) + ((0x9e3779b97f4a7c15 * uint64(i+1)) & (1<<64 - 1))
		z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
		z = (z ^ (z >> 27)) * 0x94d049bb133111eb
		g.state[i] = z ^ (z >> 31)
	}
}

// rotl rotates x to the left by k bit positions.
func rotl(x uint64, k uint) uint64 {
	return (x << k) | (x >> (64 - k))
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (g *PRNG) Uint64() uint64 {
	s0, s1 := g.state[0], g.state[1]
	result := s0 + s1

	s1 ^= s0
	g.state[0] = rotl(s0, 55) ^ s1 ^ (s1 << 14)
	g.state[1] = rotl(s1, 36)

	return result
}

// Int63 returns a non-negative pseudorandom 63-bit integer as an int64.
func (g *PRNG) Int63() int64 {
	return int64(g.Uint64() & (1<<63 - 1))
}
