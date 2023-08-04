package main

import "fmt"

// int 타입의 가변 매개변수를 사용하는 foo식별자로 func를 만드시오
// func에 slice 타입의 값을 전달하시오 즉 int의 슬라이스를 unfurl 하시오

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
