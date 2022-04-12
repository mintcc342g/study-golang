package callback

import (
	"fmt"
)

// GetNameResult ...
type GetNameResult func(name string) string // 타입이 함수인 커스텀 타입을 만들어줌.

// SayHello ...
func SayHello(name string, nameFn GetNameResult) { // 만든 커스텀 타입을 받아줌.
	fmt.Printf("Hello, %s", nameFn(name)) // 실제 함수로서 사용해줌.
}

func TestFunction() {
	// 이렇게 콜백함수를 파라미터로 넘김.
	// 근데 그냥 파리미터에 함수 쓴다고 해서 되는 게 아님.
	// SayHello() 함수의 파라미터에 저 콜백함수를 받을 수 있게끔 해줘야 되기 때문에 GetNameResult 라는 커스텀 타입을 만들어서 사용하고 있음.
	SayHello("World!", func(name string) string {
		return name // result: Hello, World!
	})
	// 순서는 SayHello 함수가 먼저 작동을 하고, SayHello 안으로 들어가서 nameFn()이 작동을 함. 그러면 인자로 넘겨준 저 함수 내용이 돌게 됨.
}
