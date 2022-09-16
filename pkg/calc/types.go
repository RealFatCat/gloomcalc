package calc

// Answer contains info with results of number calculations
type Answer struct {
	perc  float32
	long  uint8
	short uint8
}

// Long shows how many rounds merchant will stay only with long rests
func (a *Answer) Long() uint8 {
	return a.long
}

// Short shows how many rounds merchant will stay only with short rests
func (a *Answer) Short() uint8 {
	return a.short
}

// Perc shows difference between Long and Short in percentes
func (a *Answer) Perc() float32 {
	return a.perc
}
