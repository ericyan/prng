package splitmix64

import "testing"

func TestSplitMix64(t *testing.T) {
	cases := []struct {
		seed int64
		nums []uint64
	}{
		{
			0,
			[]uint64{
				16294208416658607535,
				7960286522194355700,
				487617019471545679,
				17909611376780542444,
				1961750202426094747,
				6038094601263162090,
				3207296026000306913,
				14232521865600346940,
				4532161160992623299,
				17561866513979060390,
			},
		},
		{
			1,
			[]uint64{
				10451216379200822465,
				13757245211066428519,
				17911839290282890590,
				8196980753821780235,
				8195237237126968761,
				14072917602864530048,
				16184226688143867045,
				9648886400068060533,
				5266705631892356520,
				14646652180046636950,
			},
		},
		{
			0xdeadbeef,
			[]uint64{
				0x4adfb90f68c9eb9b,
				0xde586a3141a10922,
				0x21fbc2f8e1cfc1d,
				0x7466ce737be16790,
				0x3bfa8764f685bd1c,
				0xab203e503cb55b3f,
				0x5a2fdc2bf68cedb3,
				0xb30a4ccf430b1b5a,
				0xa90415039bd5985,
				0x26ae50847745eb7e,
			},
		},
		{
			0x0ddc0ffeebadf00d,
			[]uint64{
				0x916df851e2aee44,
				0x9ade0f09ffca1bc4,
				0xea0f3dbc1648bdab,
				0x534893728891ef92,
				0x9dc9a59d9c4bfe18,
				0xf1d11238721897a4,
				0xa53d985a137009ac,
				0xdf2efc9dc6e3dcbf,
				0x43b655792e001cdf,
				0x57373124fb6b7557,
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
