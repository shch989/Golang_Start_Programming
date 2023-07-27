package main

import "fmt"

// 모든 대문자 알파벳에 대해 룬 코드 포인트를 세번씩 출력하시오

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
