package interfaces

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"testing"
)

// TestImgFormat ...
type TestImgFormat struct { // 테스트 용도의 구조체 선언
	file string // 로컬 디스크에 있는 파일 디렉토리 위치를 명시해주기 위한 필드
}

// Open ...테스트 구조체에 MyImg interface에 명시한 메소드를 구현해줌. 이건 실제 디스크에 있는 이미지 파일을 열기 위한 용도임.
func (tif *TestImgFormat) Open() (multipart.File, error) {
	return os.Open(tif.file) // 그래서 실제 내용물은 os.Open() 을 사용해줌.
}

func TestCheckImg(t *testing.T) {
	// 테스트를 위한 구조체 선언
	type args struct {
		file MyImg
	}

	// 테스트에 사용할 이미지 파일 구조체. 이걸 args 구조체의 file 필드에 넣을거임.
	ok := &TestImgFormat{file: "./cute_404_docker.png"}
	nok := &TestImgFormat{file: "./checkingImgFormat.go"}

	// 테스트셋을 만듦.
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{file: ok}, ""},
		{"t2", args{file: nok}, "unknown"}, // CheckImg 함수는 error 를 리턴하기 때문에, 만약 에러를 검사하려면 리턴하는 에러의 문자열을 비교해야함..
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := CheckImg(test.args.file); !strings.Contains(fmt.Sprintf("%v", got), test.want) {
				t.Errorf("CheckImg() = %v", got)
			}
		})
	}
}
