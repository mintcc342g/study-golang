package pointers

import "fmt"

// BakeBiscuitFailureVersion ...
func BakeBiscuitFailureVersion() {
	// 비스킷 만들기

	// 초코 비스킷을 만들거임.
	chocoBiscuit := &Biscuit{}

	// 2. 재료를 한꺼번에 넣어주자 (초코렛, 밀가루, 소금, 우유)
	chocoBiscuit.Choco.Cacao = 5
	chocoBiscuit.Choco.Milk = "프랑스산 우유 50g"
	chocoBiscuit.Choco.Sugar = 30
	chocoBiscuit.Flour = "스위스산 밀가루 100g"
	chocoBiscuit.Salt = 12
	chocoBiscuit.Milk = 50

	// 3. 만든 결과는..?
	fmt.Printf("맛있는 초코 비스켓 완성! >>> %v", chocoBiscuit) // exit status 2 -> 즉 대실패!

	/*
		// 왜 에러가 났을까?
			- 초콜릿의 재료를 담아줄 메모리 공간을 할당해주지 않았기 때문임. (누가 요리 만들 때 재료 한꺼번에 넣으래..?)
			- 즉, 포인터 문제

		// 포인터 (개인적으로 이해한..)
			- Biscuit 구조체는 내부 필드인 Choco가 Chocolate 구조체의 포인터 타입임.
				- 풀어서 말하자면, Choco 필드는 Chocolate 구조체의 메모리 주소를 가리키고 있다는 것이 됨.
				- 근데 이게 메모리 주소를 가리킨다고만 했지, 실제 메모리 내의 특정 주소를 가리킨다는 뜻은 아님.
			- 즉, Choco 필드는 "나는 Chocolate 구조체의 메모리 주소 가리킬거야~ (아직 그 주소를 모르지만..)" 라고 말하고 있는 것

			- 포인터는 메모리 주소를 가리킨다는데, 왜 주소를 모르는 걸까?
				- Choco의 Chocolate 구조체 포인터 타입은 그냥 타입으로서 선언되어 있을 뿐, Chocolate 구조체가 실제 어떤 변수에 할당되기 전 까지는 메모리에 올라간 건 아님.
				- 객체는 인스턴스화 됐을 때 메모리에 올라가니까. 그래서 진짜 메모리 주소는 모르는 상태.
				- 그래서 "메모리 주소를 가리킬거야~ (아직 그 주소는 모르지만.. 인스턴스화 되면 알겠지?)" 라는 것

			- 정리하자면..
				- * 은 메모리 주소를 가리킬 것이라는 일종의 암시? 예고? 선언? 과도 같음. 실제 메모리 주소를 가리키진 않음.
				- & 야말로 실제 메모리 주소를 가리킴. 정확히는, 그 메모리 공간을 만들어주는 역할을 하는 것 같음. 지가 만들었으니까 그 주소도 당연히 알겠지...?

		// 따라서, 위의 예제에서도 chocoBiscuit 에 선언된 Biscuit 객체 내부에 Choco 필드를 위한 Chocoate 객체를 함께 선언해서 인스턴스화 해주어야 정상적으로 돌아감.
	*/
}

// BakeBiscuitSuccessVersion ...
func BakeBiscuitSuccessVersion() {

	// 다시 초코 비스킷을 만들어보자
	chocoBiscuitRetry := &Biscuit{Choco: &Chocolate{}} // 이렇게 미리 초콜렛을 만들기 위한 별도의 그릇을 만들어주자

	// 2. 다시 재료를 넣자.
	// 초코렛을 위한 그릇에 초코렛을 만들어주고..
	chocoBiscuitRetry.Choco.Cacao = 5
	chocoBiscuitRetry.Choco.Milk = "프랑스산 우유 50g"
	chocoBiscuitRetry.Choco.Sugar = 30

	// 비스킷을 만들어주자
	chocoBiscuitRetry.Flour = "스위스산 밀가루 100g"
	chocoBiscuitRetry.Salt = 12
	chocoBiscuitRetry.Milk = 50

	// 3. 만든 결과는..?
	fmt.Printf("맛있는 초코 비스킷 완성! >>> %v", chocoBiscuitRetry) // 맛있는 초코 비스킷 완성! [&{0xc0000040a0 스위스산 밀가루 100g 12 50}]
}
