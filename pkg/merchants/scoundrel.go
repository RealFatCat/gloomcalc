package merchants

const (
	scoundrelHandSize = 9
)

// NewScoundrel returns *Merchant for Scoundrel
func NewScoundrel() *Merchant {
	return &Merchant{
		handSize: scoundrelHandSize,
	}
}
