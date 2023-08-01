package main

import "fmt"

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("해당 배열의 숫자들의 합은 ", x, " 입니다.")
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
