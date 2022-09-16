package merchants

const (
	spellweaverHandSize = 8
)

// NewSpellweaver returns *Merchant for Spellweaver
func NewSpellweaver() *Merchant {
	return &Merchant{
		handSize:             spellweaverHandSize,
		restoreBurnedAbility: true,
	}
}
