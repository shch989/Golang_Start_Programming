package main

import "fmt"

func main() {
	if x := 42; x == 42 {
		fmt.Println(x)
	}
	// 에러 발생
	// fmt.Println(x)
}
