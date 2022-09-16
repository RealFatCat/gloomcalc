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
	s, l = calc(m.HandSize(), m.RestoreBurnedAbility())
	return
}

func calc(hs uint8, ub bool) (s, l uint8) {
	var f uint8                // number of rounds before first rest
	for i := hs; i >= 2; i-- { // stop, when it just one card on hand left
		s += i / 2 // player must use two cards, so first rest will be after the "number of cards in hand divided by two" round
		if ub && i == hs {
			f = s // save first result here, we need it only if merchant have "restore all burned cards" card
		}
		if i > 2 {
			l++ // long rest gives one more round to play
		}
	}

	if ub { // if merchant can restore it's burned cards
		s = 2*s - f
		l = 2*l - 1
	}
	return
}
