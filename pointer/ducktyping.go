package pointer

import "fmt"

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

func TestDuckTyping() {
	ice := &IceCream{
		IceType: "banila ice cream",
	}

	ice.Order("affogato")
	fmt.Printf("\n\nmy ice cream >> %s", ice.GetMyIceCream()) // my ice cream: banila ice cream with ... espresso

	ice.Order("chocolate")
	fmt.Printf("\n\nmy ice cream >> %s", ice.GetMyIceCream()) // my ice cream: banila ice cream with ... chocolate syrup

	ice.Order("steak")
	fmt.Printf("\n\nmy ice cream >> %s", ice.GetMyIceCream()) // my ice cream: banila ice cream with ... What?
}
