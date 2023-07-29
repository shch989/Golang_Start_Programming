package main

import "fmt"

// 합성 리터럴을 사용하여 정수형 값 5개를 갖는 배열을 만드시오
// 각각의 인덱스 위치에 값을 할당하여라

func main() {
	x := [5]int{42, 43, 44, 45, 46}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
