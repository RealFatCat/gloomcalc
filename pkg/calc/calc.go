package calc

import (
	"github.com/realfatcat/gloomcalc/pkg/merchants"
)

// Rounds calculates number of rounds for merchant
func Rounds(m *merchants.Merchant) *Answer {
	s, l := calcRounds(m)
	p := (1 - float32(s)/float32(l+s)) * 100
	return &Answer{perc: p, long: s + l, short: s}
}

func calcRounds(m merchants.Merchanter) (s, l uint8) {
	s, l = calc(m.HandSize())
	if m.RestoreBurnedAbility() {
		st, lt := calc(m.HandSize() - 1)
		s += st
		l += lt
	}
	return
}

func calc(hs uint8) (s, l uint8) {
	for i := hs; i >= 2; i-- {
		s += i / 2
		if i > 2 {
			l++
		}
	}
	return
}
