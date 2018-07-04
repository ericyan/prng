package prng

import (
	"math/rand"

	"github.com/ericyan/prng/internal/splitmix64"
)

// A xoroshiro128Plus is an instance of xoroshiro128+ generator.
type xoroshiro128Plus [2]uint64

// NewXoroshiro128Plus returns a new xoroshiro128+ generator seeded with
// the given value.
func NewXoroshiro128Plus(seed int64) rand.Source64 {
	s := new(xoroshiro128Plus)
	s.Seed(seed)
	return s
}

// Seed initializes the generator to a deterministic state.
//
// It first seeds a SplitMix64 generator with the given 64-bit value and
// then uses the outputs to seed the xoroshiro128+ generator which has
// 128 bits of state.
func (s *xoroshiro128Plus) Seed(seed int64) {
	r := splitmix64.New(seed)
	s[0], s[1] = r.Uint64(), r.Uint64()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (s *xoroshiro128Plus) Uint64() uint64 {
	s0, s1 := s[0], s[1]
	result := s0 + s1

	s1 ^= s0
	s[0] = rotl(s0, 24) ^ s1 ^ (s1 << 16)
	s[1] = rotl(s1, 37)

	return result
}

// Int63 returns a non-negative pseudorandom 63-bit integer as an int64.
func (s *xoroshiro128Plus) Int63() int64 {
	return int63(s.Uint64())
}

// A xoroshiro128StarStar is an instance of xoroshiro128** generator.
type xoroshiro128StarStar [2]uint64

// NewXoroshiro128StarStar returns a new xoroshiro128** generator seeded
// with the given value.
func NewXoroshiro128StarStar(seed int64) rand.Source64 {
	s := new(xoroshiro128StarStar)
	s.Seed(seed)
	return s
}

// Seed initializes the generator to a deterministic state.
//
// It first seeds a SplitMix64 generator with the given 64-bit value and
// then uses the outputs to seed the xoroshiro128** generator which has
// 128 bits of state.
func (s *xoroshiro128StarStar) Seed(seed int64) {
	r := splitmix64.New(seed)
	s[0], s[1] = r.Uint64(), r.Uint64()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (s *xoroshiro128StarStar) Uint64() uint64 {
	s0, s1 := s[0], s[1]
	result := rotl(s0*5, 7) * 9

	s1 ^= s0
	s[0] = rotl(s0, 24) ^ s1 ^ (s1 << 16)
	s[1] = rotl(s1, 37)

	return result
}

// Int63 returns a non-negative pseudorandom 63-bit integer as an int64.
func (s *xoroshiro128StarStar) Int63() int64 {
	return int63(s.Uint64())
}
