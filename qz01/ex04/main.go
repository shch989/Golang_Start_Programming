package main

import "fmt"

// int 타입을 가진 커스텀 타입을 하나 만든다.
// var 키워드를 사용해서 식별자가 x인 커스텀 타입의 변수를 만든다.
// main 함수에서 변수 x의 값과 x의 타입을 프린트한다.
// = 연산자를 사용해서 변수 x에 42를 지정한 하 다시 프린트한다.

type galaxy int

var x galaxy

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
