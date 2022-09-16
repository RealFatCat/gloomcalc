package merchants

const (
	craigheartHandSize = 11
)

// NewCraigheart returns *Merchant for Craigheart
func NewCraigheart() *Merchant {
	return &Merchant{
		handSize: craigheartHandSize,
	}
}
