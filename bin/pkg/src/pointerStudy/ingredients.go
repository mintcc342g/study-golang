package pointerstudy

// Biscuit ...
type Biscuit struct {
	Choco *Chocolate
	Flour string
	Salt  uint64
	Milk  uint64
}

// Chocolate ...
type Chocolate struct {
	Cacao int64
	Milk  string
	Sugar float64
}
