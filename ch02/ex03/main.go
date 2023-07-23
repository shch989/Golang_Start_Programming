package main

import "fmt"

func main() {
	s := "Hello playground"
	fmt.Println(s)        // "Hello playground" 문자열 출력
	fmt.Printf("%T\n", s) // 문자열 변수의 타입 출력

	bs := []byte(s)        // 문자열을 바이트 슬라이스로 변환
	fmt.Println(bs)        // 바이트 슬라이스 출력
	fmt.Printf("%T\n", bs) // 바이트 슬라이스의 타입 출력

	for i := 0; i < len(s); i++ { // 문자열의 각 문자를 UTF-8 형식으로 출력
		fmt.Printf("%#U\n", s[i]) // 각 문자를 UTF-8 형식으로 출력
	}

	for i, v := range s { // 문자열의 각 문자를 인덱스와 함께 출력
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
