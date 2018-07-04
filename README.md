# PRNG

A pure Go implementation of the xoshiro/xoroshiro family of pseudorandom
number generators. Compared with the default additive lagged fibonacci
generator in the `math/rand` package, these generators are reportedly
faster and produce superior statistical results.

The algorithms are designed by Sebastiano Vigna in collaboration with
David Blackman. More details about the generators can be found at
http://xoroshiro.di.unimi.it/

## xoshiro256** and xoshiro256+

xoshiro256** (XOR/shift/rotate) is a all-purpose, rock-solid generator.
It has excellent speed, a state of 256 bits that is large enough for any
parallel application, and it passes all tests we are aware of.

xoshiro256+ is a variant optimised for generating only 64-bit floating
point numbers (by extracting the upper 53 bits). It passes all the tests
except for the lowest three bits, which might fail linearity tests (and
just those).

## xoroshiro128** and xoroshiro128+

xoroshiro128** (XOR/rotate/shift/rotate) and xoroshiro128+ has 128 bits
of state. They are suitable only for low-scale parallel applications.

It is worth noting that xoroshiro128+ exhibits a mild dependency in
Hamming weights that generates a failure after 8 TB of output in a test.
However, this slight bias should not affect any application.

## Limitations

 * Not safe for concurrent use by multiple goroutines.
 * Not cryptographically secure.
