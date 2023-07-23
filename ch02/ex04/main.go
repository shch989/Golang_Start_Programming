package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s) // "H" 문자열 출력

	bs := []byte(s) // 문자열을 바이트 슬라이스로 변환
	fmt.Println(bs) // 바이트 슬라이스 출력

	n := bs[0]             // 바이트 슬라이스의 첫 번째 요소를 정수로 할당
	fmt.Println(n)         // 정수 값 출력
	fmt.Printf("%T\n", n)  // 정수 값의 타입 출력
	fmt.Printf("%b\n", n)  // 정수 값을 2진수로 출력
	fmt.Printf("%#X\n", n) // 정수 값을 16진수로 출력
}
