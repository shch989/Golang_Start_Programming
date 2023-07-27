package main

import "fmt"

// 이전과 다른 반복문으로 현재까지 살아온 년도들을 출력하라

func main() {
	bd := 1998
	for {
		if bd > 2023 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
