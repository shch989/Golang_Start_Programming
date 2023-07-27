package main

import "fmt"

// 10부터 100까지의 수들을 4로 나눈 나머지 값을 출력하라

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v에서 4를 나눌 경우 나머지 값은 %v가 나옵니다.\n", i, i%4)
	}
}
