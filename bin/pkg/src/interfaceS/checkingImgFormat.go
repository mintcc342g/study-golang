package interfaces

import (
	"image" // DecodeConfig() 를 사용하기 위해 image 패키지 import
	"mime/multipart"

	// 이미지로 허용해줄 확장자를 패키지로 추가함.
	// 블랭크 임포트를 사용하는데, main 패키지가 아니면 주석을 달아줘야 lint 에러가 안 남.
	_ "image/jpeg"
	// png
	_ "image/png"
)

/*
func CheckImg(src *multipart.FileHeader) error { // src로 들어온 이미지 파일이 실제 이미지 확장자를 가진 파일인지 아닌지를 검사하는 유틸성 함수
	r, err := src.Open() // Open을 해줘야 io.Reader 타입의 데이터가 나옴.
	if err != nil {
		return err
	}
	// DecodeConfig() 를 거칠 때 블랭크 임포트를 한 패키지들의 init이 실행됨. 거기서 확장자 검사가 되는 것.
	// 참고로,  리턴값은 이미지config, 파일foramt, err 의 3개
	_, _, err = image.DecodeConfig(r)
	if err != nil {
		return err
	}

	return nil

}

// 처음에는 이런 식으로 짰는데..
	- 이 함수만을 위한 테스트코드를 짜려니까 코드 짜는 게 불가능한 것 같았음.
	- multipart.FileHeader는 http request의 form-data 형식으로 들어온 데이터를 파싱할 때 얻을 수 있고, FileHeader 객체를 따로 만드는 것은 불가능했음.
	- 그렇다고 이 함수만을 위한 api를 만들수도 없는 노릇이고.
	- 게다가 실제 테스트를 하려면 이미지 파일을 열어야 하는데, 그건 바로 io.Reader 타입으로 나와버린다는 문제도 있었음.
	- 그 때 시니어분이 알려준 방법이 interface의 활용!

// go의 interface는 동일한 메소드를 구현하고 있으면 되는 규약? 이라는 듯
	- 어떤 구조체가 interface에 구현된 메소드를 갖고 있기만 하면, 해당 interface 타입으로 인정을 받는 것 같음.
*/

// MyImg ...
type MyImg interface { // 이렇게 inerface를 선언하고
	Open() (multipart.File, error) // multipart.FileHeader 가 갖고 있는 Open() 메소드랑 동일한 메소드를 인터페이스에 선언해줌.
}

// CheckImg ...
func CheckImg(src MyImg) error { // 이미지 검사 함수의 파라미터로 인터페이스를 넣어주자
	// multipart.FileHeader struct는 MyImg 인터페이스가 갖고 있는 Open() 메소드를 구현하고 있는 것으로 인지됨.
	// 그래서 파라미터로 무사히 들어오고, multipart.FileHeader의 Open() 메소드가 사용되어 파일이 무사히 읽히게 된다!
	// 테스트코드도 이런 방식으로 해결을 했다. checkingImgFormat_test.go 파일을 확인해보면..
	r, err := src.Open()
	if err != nil {
		return err
	}
	_, _, err = image.DecodeConfig(r)
	if err != nil {
		return err
	}

	return nil
}
