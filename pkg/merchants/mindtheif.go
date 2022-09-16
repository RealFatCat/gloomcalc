package merchants

const (
	mindthiefHandSize = 10
)

// NewMindthief returns *Merchant for Mindthief
func NewMindthief() *Merchant {
	return &Merchant{
		handSize: mindthiefHandSize,
	}
}
