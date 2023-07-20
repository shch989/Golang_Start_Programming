package main

import "fmt"

// var 키워드를 사용해서 식별자가 y인 변수를 만든다.
// 변수의 실제 타입은 커스텀 타입 x여야 한다.
// 변환을 이용해서 x에 저장된 값의 타입을 실제 타입으로 변환한다.
// 단축 선언 연산자를 사용해서 그 값을 y에 지정하고 y에 저장된 값을 프린트한다.
// y의 타입을 프린트한다.

type galaxy int

var x galaxy
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
