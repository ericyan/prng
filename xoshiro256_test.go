package prng

import (
	"testing"
)

func TestXoshiro256Plus(t *testing.T) {
	cases := []struct {
		seed [4]uint64
		nums []uint64
	}{
		{
			[4]uint64{1, 2, 3, 4},
			[]uint64{
				5,
				211106232532999,
				211106635186183,
				9223759065350669058,
				9250833439874351877,
				13862484359527728515,
				2346507365006083650,
				1168864526675804870,
				34095955243042024,
				3466914240207415127,
			},
		},
		{
			[4]uint64{0x916df851e2aee44, 0x9ade0f09ffca1bc4, 0xea0f3dbc1648bdab, 0x534893728891ef92},
			[]uint64{
				0x5c5f72f7a6bcddd6,
				0x3f0b1d313d0088fd,
				0xaca578e5f3d68156,
				0xd00318890156ba76,
				0x921bc790fa6e212a,
				0x7e43dd7d95155fcb,
				0x1994a83e2b874bc,
				0x2116e8eb1edb4347,
				0xc586f85ae89d4a46,
				0x3269586c32a40edc,
			},
		},
	}

	for _, c := range cases {
		s := new(xoshiro256Plus)
		s[0], s[1], s[2], s[3] = c.seed[0], c.seed[1], c.seed[2], c.seed[3]

		for i, n := range c.nums {
			if got := s.Uint64(); got != n {
				t.Errorf("xoshiro256+{%d,%d,%d,%d}[%d]: expect %d, got %d\n", c.seed[0], c.seed[1], c.seed[2], c.seed[3], i, n, got)
			}
		}
	}
}

func TestXoshiro256StarStar(t *testing.T) {
	cases := []struct {
		seed [4]uint64
		nums []uint64
	}{
		{
			[4]uint64{1, 2, 3, 4},
			[]uint64{
				11520,
				0,
				1509978240,
				1215971899390074240,
				1216172134540287360,
				607988272756665600,
				16172922978634559625,
				8476171486693032832,
				10595114339597558777,
				2904607092377533576,
			},
		},
		{
			[4]uint64{0x916df851e2aee44, 0x9ade0f09ffca1bc4, 0xea0f3dbc1648bdab, 0x534893728891ef92},
			[]uint64{
				0x845260fb4370ba1b,
				0x1258cdc44a57c930,
				0xd44441aafba9507a,
				0x71d639bfce47e6d3,
				0xf9b2afeec393909e,
				0xd25513e17c0ce113,
				0xfc2f023b3464c57b,
				0xcfa56a15950b5b43,
				0xea666f369aed503c,
				0xff7e25ec8a3ab0e0,
			},
		},
	}

	for _, c := range cases {
		s := new(xoshiro256StarStar)
		s[0], s[1], s[2], s[3] = c.seed[0], c.seed[1], c.seed[2], c.seed[3]

		for i, n := range c.nums {
			if got := s.Uint64(); got != n {
				t.Errorf("xoshiro256**{%d,%d,%d,%d}[%d]: expect %d, got %d\n", c.seed[0], c.seed[1], c.seed[2], c.seed[3], i, n, got)
			}
		}
	}
}
