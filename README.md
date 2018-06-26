# PRNG

A Go implementation of the xoroshiro128+ pseudorandom number generator.
It is reportedly faster and produces superior statistical results to the
default additive lagged fibonacci generator in the `math/rand` package.

## xoroshiro128+

xoroshiro128+ (XOR/rotate/shift/rotate) is currently the fastest full
cycle pseudorandom number generator passing the Big Crush tests without
systematic failures.

It is designed by Sebastiano Vigna in collaboration with David Blackman.
More information is available at http://xoroshiro.di.unimi.it/

## Limitations

 * Not safe for concurrent use by multiple goroutines.
 * Not cryptographically secure.
