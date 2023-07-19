package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// a와 b는 같은 int 값을 가지고 있으나 타입의 경우
	// a는 int b는 hotdog라는 이름으로 지정하였기 때문에
	// 타입 에러가 발생한다
	// a = b
}
