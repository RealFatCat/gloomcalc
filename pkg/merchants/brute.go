package merchants

const (
	bruteHandSize = 10
)

// NewBrute returns *Merchant for Brute
func NewBrute() *Merchant {
	return &Merchant{
		handSize: bruteHandSize,
	}
}
