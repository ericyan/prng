package prng

import (
	"testing"
)

func TestXoroshiro128Plus(t *testing.T) {
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
		{
			[2]uint64{0x916df851e2aee44, 0x9ade0f09ffca1bc4},
			[]uint64{
				11814330020949985800,
				11817088786836023749,
				1654166990350674155,
				14112748191344281834,
				4288295283113472773,
				8391955421631067594,
				168274855724945977,
				2815117763357611551,
				12187186948608395331,
				10629044371437376348,
			},
		},
	}

	for _, c := range cases {
		s := NewXoroshiro128Plus(c.seed)
		for i, n := range c.nums {
			if got := s.Uint64(); got != n {
				t.Errorf("Xoroshiro128PlusPlus{%d,%d}[%d]: expect %d, got %d\n", c.seed[0], c.seed[1], i, n, got)
			}
		}
	}
}
