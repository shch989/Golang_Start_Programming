package main

import "fmt"

// 이전 예제에서 else if와 else를 사용하여 프로그램을 만드시오

func main() {
	x := "James Bond"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDBONDBONDBOND", x)
	} else {
		fmt.Println("neither")
	}
}
