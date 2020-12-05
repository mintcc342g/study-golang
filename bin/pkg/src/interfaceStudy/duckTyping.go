package interfacestudy

// IceCream ...
type IceCream struct {
	IceType string
	Topping Topping
}

// Order ...
func (i *IceCream) Order(name string) {
	t, ok := map[string]Topping{
		"chocolate": &ChocolateIce{},
		"affogato":  &Affogato{},
	}[name]

	if ok {
		i.Topping = t
	} else {
		i.Topping = &BadRequest{}
	}
}

// GetMyIceCream ...
func (i *IceCream) GetMyIceCream() string {
	return i.IceType + " with ... " + i.Topping.GetTopping()
}

// Topping for IceCream
type Topping interface {
	GetTopping() string
}

// ChocolateIce ...
type ChocolateIce struct{}

// GetTopping ... Topping interface
func (c *ChocolateIce) GetTopping() string {
	return "chocolate syrup"
}

// Affogato ...
type Affogato struct{}

// GetTopping ... Topping interface
func (a *Affogato) GetTopping() string {
	return "espresso"
}

// BadRequest ...
type BadRequest struct{}

// GetTopping ... Topping interface
func (b *BadRequest) GetTopping() string {
	return "What?"
}
