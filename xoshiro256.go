// Package prng implements xoroshiro128+ pseudorandom number generators.
package prng

import (
	"math/rand"

	"github.com/ericyan/prng/internal/splitmix64"
)

// A xoshiro256Plus is an instance of xoshiro256+ generator.
type xoshiro256Plus [4]uint64

// NewXoroshiro128Plus returns a new xoshiro256+ generator seeded with
// the given value.
func NewXoshiro256Plus(seed int64) rand.Source64 {
	s := new(xoshiro256Plus)
	s.Seed(seed)
	return s
}

// Seed initializes the generator to a deterministic state.
//
// It first seeds a SplitMix64 generator with the given 64-bit value and
// then uses the outputs to seed the xoshiro256+ generator which has 256
// bits of state.
func (s *xoshiro256Plus) Seed(seed int64) {
	r := splitmix64.New(seed)
	s[0], s[1], s[2], s[3] = r.Uint64(), r.Uint64(), r.Uint64(), r.Uint64()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (s *xoshiro256Plus) Uint64() uint64 {
	result := s[0] + s[3]

	t := s[1] << 17

	s[2] ^= s[0]
	s[3] ^= s[1]
	s[1] ^= s[2]
	s[0] ^= s[3]

	s[2] ^= t

	s[3] = rotl(s[3], 45)

	return result
}

// Int63 returns a non-negative pseudorandom 63-bit integer as an int64.
func (s *xoshiro256Plus) Int63() int64 {
	return int63(s.Uint64())
}

// A  xoshiro256StarStar is an instance of xoshiro256** generator.
type xoshiro256StarStar [4]uint64

// NewXoroshiro128StarStar returns a new xoshiro256** generator seeded
// with the given value.
func NewXoshiro256StarStar(seed int64) rand.Source64 {
	s := new(xoshiro256StarStar)
	s.Seed(seed)
	return s
}

// Seed initializes the generator to a deterministic state.
//
// It first seeds a SplitMix64 generator with the given 64-bit value and
// then uses the outputs to seed the xoshiro256** generator which has
// 256 bits of state.
func (s *xoshiro256StarStar) Seed(seed int64) {
	r := splitmix64.New(seed)
	s[0], s[1] = r.Uint64(), r.Uint64()
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (s *xoshiro256StarStar) Uint64() uint64 {
	result := rotl(s[1]*5, 7) * 9

	t := s[1] << 17

	s[2] ^= s[0]
	s[3] ^= s[1]
	s[1] ^= s[2]
	s[0] ^= s[3]

	s[2] ^= t

	s[3] = rotl(s[3], 45)

	return result
}

// Int63 returns a non-negative pseudorandom 63-bit integer as an int64.
func (s *xoshiro256StarStar) Int63() int64 {
	return int63(s.Uint64())
}
