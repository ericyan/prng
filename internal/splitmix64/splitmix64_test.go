package splitmix64

import "testing"

func TestSplitMix64(t *testing.T) {
	cases := []struct {
		seed int64
		nums []uint64
	}{
		{
			42,
			[]uint64{
				13679457532755275413,
				2949826092126892291,
				5139283748462763858,
				6349198060258255764,
				701532786141963250,
				16015981125662989062,
				4028864712777624925,
				14769051326987775908,
				6270620877612482005,
				11408980392250668974,
			},
		},
		{
			0xdeadbeef,
			[]uint64{
				0x4adfb90f68c9eb9b,
				0xde586a3141a10922,
			},
		},
		{
			0x0ddc0ffeebadf00d,
			[]uint64{
				0x916df851e2aee44,
				0x9ade0f09ffca1bc4,
			},
		},
	}
	for _, c := range cases {
		r := New(c.seed)
		for i, n := range c.nums {
			if got := r.Uint64(); got != n {
				t.Errorf("SplitMix64{%d}[%d]: expect %d, got %d\n", c.seed, i, n, got)
			}
		}
	}
}
