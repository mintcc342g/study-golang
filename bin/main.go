package main

import (
	"fmt"

	cfs "./pkg/src/callbackFnStudy"
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

	// Study of Callback Function
	cfs.SayHello("World!", func(name string) string {
		// 이렇게 콜백함수를 파라미터로 넘김.
		// 근데 그냥 파리미터에 함수 쓴다고 해서 되는 게 아님.
		// SayHello() 함수의 파라미터에 저 콜백함수를 받을 수 있게끔 해줘야 되기 때문에 GetNameResult 라는 커스텀 타입을 만들어서 사용하고 있음.
		return name // result: Hello, World!
	})
	// 순서는 SayHello 함수가 먼저 작동을 하고, SayHello 안으로 들어가서 nameFn()이 작동을 함. 그러면 인자로 넘겨준 저 함수 내용이 돌게 됨.
}
