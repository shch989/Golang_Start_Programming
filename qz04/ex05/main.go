package main

import "fmt"

// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}를 만드시오
// x 슬라이스에서 [42, 43, 44, 48, 49, 50, 51]가 출력되도록 만드시오

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
