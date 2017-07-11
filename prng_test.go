package prng

import "testing"

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
