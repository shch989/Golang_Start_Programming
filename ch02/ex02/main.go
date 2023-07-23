package main

import (
	"fmt"
	"runtime"
)

// uint8: 8비트 부호 없는 정수 (0에서 255까지)
// uint16: 16비트 부호 없는 정수 (0에서 65535까지)
// uint32: 32비트 부호 없는 정수 (0에서 약 42억까지)
// uint64: 64비트 부호 없는 정수 (매우 큰 양의 정수 값을 저장 가능)

// int8: 8비트 부호 있는 정수 (-128에서 127까지)
// int16: 16비트 부호 있는 정수 (-32768에서 32767까지)
// int32: 32비트 부호 있는 정수 (약 -21억에서 21억까지)
// int64: 64비트 부호 있는 정수 (매우 큰 범위의 정수 값을 저장 가능)

// float32: 32비트 단정밀도 부동 소수점 숫자 (약 7자리의 유효 숫자)
// float64: 64비트 배정밀도 부동 소수점 숫자 (약 15자리의 유효 숫자)

// complex64: 32비트 실수와 허수로 구성된 복소수
// complex128: 64비트 실수와 허수로 구성된 복소수

// byte: uint8과 동일, 부호 없는 8비트 정수 (주로 문자열과 바이트 슬라이스에서 사용)
// rune: int32와 동일, 유니코드 코드 포인트를 나타냄 (주로 유니코드 문자 처리에 사용)

var x int
var y float64

func main() {
	x := 42
	y := 42.345345
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Printf(runtime.GOOS)
	fmt.Printf(runtime.GOARCH)
}
