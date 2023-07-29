package main

import "fmt"

// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}를 만들어라
// x 배열에 52를 추가하고 출력하여라
// 53 54 55 이 세 개를 x에 추가한 뒤 각 구문으로 출력하여라
// y := []int{56, 57, 58, 59, 60}를 만들어라
// x 배열에 y배열을 추가한 뒤 해당 슬라이스를 출력하여라

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	for i := 11; i <= 13; i++ {
		println(x[i])
	}

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}
