package main

import (
	"fmt"

	ins "./pkg/src/interfaceStudy"
	pos "./pkg/src/pointerStudy"
)

func main() {
	// Study of Pointer
	pos.BakeBiscuitSuccessVersion()
	// pos.BakeBiscuitFailureVersion()

	// Study of Interface ... duck typing
	ice := &ins.IceCream{
		IceType: "banila ice cream",
	}

	ice.Order("affogato")
	fmt.Printf("\n\nmy ice cream >> %s", ice.GetMyIceCream())

	ice.Order("chocolate")
	fmt.Printf("\n\nmy ice cream >> %s", ice.GetMyIceCream())

	ice.Order("steak")
	fmt.Printf("\n\nmy ice cream >> %s", ice.GetMyIceCream())

}
