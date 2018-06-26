package prng

// rotl rotates x to the left by k bit positions.
func rotl(x uint64, k uint) uint64 {
	return (x << k) | (x >> (64 - k))
}

// int63 converts a uint64 to a 63-bit integer as an int64.
func int63(x uint64) int64 {
	return int64(x & (1<<63 - 1))
}
