package main

import "fmt"

// 변수에 int형 값을 할당하고 이 값을 10진수, 2진수, 16진수로 출력하라
// 해당 변수를 1비트 왼쪽으로 이동하고 그 값을 새로운 변수에 할당하여라
// 새로 할당된 값을 10진수, 2진수, 16진수로 출력하라

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}
