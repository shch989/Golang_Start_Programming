package main

import "fmt"

// 10개의 정수 값이 있는 슬라이스를 만드시오
// 슬라이스를 슬라이싱하여 특정 값들을 갖도록 하여라

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)
}
