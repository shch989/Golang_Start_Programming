package main

import "fmt"

// 반복문으로 루프를 만들고 살아온 년도들을 출력하라

func main() {
	bd := 1998
	for bd <= 2023 {
		fmt.Println(bd)
		bd++
	}
}
