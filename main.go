package main

import (
	"flag"
	"fmt"

	"github.com/realfatcat/gloomcalc/pkg/calc"
	"github.com/realfatcat/gloomcalc/pkg/merchants"
)

var (
	merchant string
	hs       uint64
	unburn   bool
)

func init() {
	flag.StringVar(&merchant, "m", "", "original merchant to calc number of rounds for")
	flag.Uint64Var(&hs, "hs", 0, "custom merchant: set hand size")
	flag.BoolVar(&unburn, "ub", false, "custom merchant: can restore burned cards")
	flag.Parse()
}

func main() {
	if merchant != "" {
		a := baseMerchant(merchant)
		printResult(merchant, a)
	}
	if hs > 0 {
		if hs < 22 {
			a := customMerchant(hs, unburn)
			printResult("Custom", a)
		} else {
			fmt.Printf("hand size should be more than 0 and less than 22\n")
		}
	}
}

func baseMerchant(m string) (a *calc.Answer) {
	switch merchant {
	case "brute":
		a = calc.Rounds(merchants.NewBrute())
	case "scoundrel":
		a = calc.Rounds(merchants.NewScoundrel())
	case "tinkerer":
		a = calc.Rounds(merchants.NewTinkerer())
	case "spellweaver":
		a = calc.Rounds(merchants.NewSpellweaver())
	case "mindthief":
		a = calc.Rounds(merchants.NewMindthief())
	case "craigheart":
		a = calc.Rounds(merchants.NewCraigheart())
	default:
		return nil
	}
	return a
}

func customMerchant(hs uint64, unburn bool) *calc.Answer {
	// "21 cards in hand ought to be enough for anybody." (c) Bill Gates, 1981
	return calc.Rounds(merchants.NewDummy(uint8(hs), unburn))
}

func printResult(merchant string, a *calc.Answer) {
	if a != nil {
		fmt.Printf("%s:\nLong Only:\t%d\nShort Only:\t%d\nPerc:\t%f\n", merchant, a.Long(), a.Short(), a.Perc())
	} else {
		fmt.Println("Unknown merchant")
	}
}
