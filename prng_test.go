package prng

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPRNG(t *testing.T) {
	cases := []struct {
		seed [2]uint64
		nums []uint64
	}{
		{
			[2]uint64{1, 2},
			[]uint64{
				3,
				36029003177443331,
				78883775479546723,
				11565523463456473958,
				4242646275387589636,
				256749404433942852,
				11014892026844973196,
				9059353499452950543,
				8597521241247625872,
				4693915028112570637,
			},
		},
	}

	for _, c := range cases {
		s := PRNG{c.seed}
		for i, n := range c.nums {
			if got := s.Uint64(); got != n {
				t.Errorf("PRNG{%d,%d}[%d]: expect %d, got %d\n", c.seed[0], c.seed[1], i, n, got)
			}
		}
	}
}

func TestSeed(t *testing.T) {
	cases := []struct {
		seed  int64
		state [2]uint64
	}{
		{0xdeadbeef, [2]uint64{0x4adfb90f68c9eb9b, 0xde586a3141a10922}},
		{0x0ddc0ffeebadf00d, [2]uint64{0x916df851e2aee44, 0x9ade0f09ffca1bc4}},
	}

	for _, c := range cases {
		s := new(PRNG)
		s.Seed(c.seed)
		if got := s.state; got != c.state {
			t.Errorf("Seed(%x): expect PRNG{%x,%x}, got PRNG{%x,%x}\n", c.seed, c.state[0], c.state[1], got[0], got[1])
		}
	}
}

func ExamplePRNG() {
	r := rand.New(new(PRNG))
	r.Seed(0x0ddc0ffeebadf00d)

	for i := 0; i < 10; i++ {
		fmt.Println(r.Uint64())
	}

	// Output:
	// 11814330020949985800
	// 11817088786836023749
	// 1654166990350674155
	// 14112748191344281834
	// 4288295283113472773
	// 8391955421631067594
	// 168274855724945977
	// 2815117763357611551
	// 12187186948608395331
	// 10629044371437376348
}

func benchmarkRandSource(s rand.Source, b *testing.B) {
	r := rand.New(s)
	r.Seed(0x0ddc0ffeebadf00d)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Uint64()
	}
	b.SetBytes(8)
}

func BenchmarkPRNG(b *testing.B) {
	benchmarkRandSource(new(PRNG), b)
}

func BenchmarkMathRand(b *testing.B) {
	benchmarkRandSource(rand.NewSource(1), b)
}
