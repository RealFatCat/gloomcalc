package merchants

// NewDummy returns *Merchant for custom merchant
func NewDummy(hs uint8, ub bool) *Merchant {
	return &Merchant{
		handSize:             hs,
		restoreBurnedAbility: ub,
	}
}
