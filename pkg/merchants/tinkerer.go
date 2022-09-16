package merchants

const (
	tinkererHandSize = 12
)

// NewTinkerer returns *Merchant for Tinkerer
func NewTinkerer() *Merchant {
	return &Merchant{
		handSize: tinkererHandSize,
	}
}
