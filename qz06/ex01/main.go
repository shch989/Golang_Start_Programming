package main

import "fmt"

// int를 반환하는 식별자 foo를 사용하여 func를 만드시오
// int와 string을 반환하는 식별자 bar로 다른 func를 만드시오
// 두 func을 모두 호출하고 결과를 출력하시오

func main() {
	n := foo()
	x, s := bar()
	fmt.Println(n, x, s)
}

func foo() int {
	return 24
}

func bar() (int, string) {
	return 1998, "조성현"
}
