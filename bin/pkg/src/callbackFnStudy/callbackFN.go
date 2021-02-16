package callbackfnstudy

import (
	"fmt"
)

// GetNameResult ...
type GetNameResult func(name string) string // 타입이 함수인 커스텀 타입을 만들어줌.

// SayHello ...
func SayHello(name string, nameFn GetNameResult) { // 만든 커스텀 타입을 받아줌.
	fmt.Printf("Hello, %s", nameFn(name)) // 실제 함수로서 사용해줌.
}
