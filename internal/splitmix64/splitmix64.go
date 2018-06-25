// Package splitmix64 implements a fast splittable pseudorandom number
// generator for initializing the state of other generators.
package splitmix64

import "math/rand"

// SplitMix64 has 64 bits of state which can be seeded with any value.
type SplitMix64 uint64

// New returns a new SplitMix64 generator seeded with the given value.
func New(seed int64) rand.Source64 {
	x := new(SplitMix64)
	x.Seed(seed)
	return x
}

// Seed uses the provided seed value to initialise the generator to a
// deterministic state.
func (x *SplitMix64) Seed(seed int64) {
	*x = SplitMix64(seed)
}

// Uint64 implements rand.Source64.
func (x *SplitMix64) Uint64() uint64 {
	*x += 0x9e3779b97f4a7c15
	z := uint64(*x)
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

// Uint64 implements rand.Source.
func (x *SplitMix64) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))
}
