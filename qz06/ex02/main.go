package main

import "fmt"

// 식별자가 foo인 다음과 같은 함수를 만드십시오.
// 타입 int인 가변 매개변수를 받음
// 타입 []int인 값을 함수에 넣은([]int를 unfurl)
// 넣어진 타입 int인 모든 값들의 합계를 리턴함
// 식별자가 bar인 다음과 같은 함수를 만드십시오.
// 타입 []int인 매개변수를 받음
// 넣어진 타입 int인 모든 값들의 합계를 리턴함

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	fmt.Println(n)

	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := bar(ii2)
	fmt.Println(n2)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
