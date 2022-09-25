package calc

import (
	"testing"

	"github.com/realfatcat/gloomcalc/pkg/merchants"
)

type tCase struct {
	want *Answer
	m    *merchants.Merchant
}

// don't want to use reflect, or external libs
func eq(a, b *Answer) bool {
	return a.perc == b.perc && a.short == b.short && a.long == b.long
}

func TestRounds(t *testing.T) {
	// if this cases works, asume that calcs works
	tCases := []tCase{
		// Brute
		{
			want: &Answer{long: 33, short: 25, perc: 24.242424},
			m:    merchants.NewBrute(),
		},
		// Scoundrel
		{
			want: &Answer{long: 27, short: 20, perc: 25.925928},
			m:    merchants.NewScoundrel(),
		},
		// Spellweaver
		{
			want: &Answer{long: 39, short: 28, perc: 28.205126},
			m:    merchants.NewSpellweaver(),
		},
		// Tinkerer
		{
			want: &Answer{long: 46, short: 36, perc: 21.739132},
			m:    merchants.NewTinkerer(),
		},
	}
	for _, tc := range tCases {
		a := Rounds(tc.m)
		if !eq(a, tc.want) {
			t.Fatalf("answer is not valid for %#v, want: %#v, got: %#v\n", tc.m, tc.want, a)
		}
	}
}
