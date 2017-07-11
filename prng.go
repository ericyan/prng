// Package prng implements xoroshiro128+ pseudorandom number generators.
package prng

// A PRNG is an instance of xoroshiro128+ generator.
type PRNG struct {
	state [2]uint64
}

// New creates and returns a new PRNG. The seed should be two unsigned
// 64-bit integers and must not be everywhere zero.
func New(seed [2]uint64) *PRNG {
	return &PRNG{seed}
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
