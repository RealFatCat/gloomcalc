package merchants

// Merchanter shows which methods merchant shoud have
type Merchanter interface {
	HandSize() uint8
	RestoreBurnedAbility() bool
}

// Merchant describes one of the merchants.
type Merchant struct {
	handSize             uint8
	restoreBurnedAbility bool
}

// HandSize returns the starter number of merchant cards
func (m *Merchant) HandSize() uint8 {
	return m.handSize
}

// RestoreBurnedAbility shows if merchant can restore all its burned cards (acutual for Spellweaver)
func (m *Merchant) RestoreBurnedAbility() bool {
	return m.restoreBurnedAbility
}
