package main

import "fmt"

// 이전 실습과 같이 합성 리터럴을 만들어서 슬라이스를 만드시오
// 단, 슬라이스에 10개의 정수 값을 주시오

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
