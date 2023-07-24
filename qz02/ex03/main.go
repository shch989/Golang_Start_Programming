package main

import "fmt"

// typed 상수와 untyped 상수를 만들어라
// 상수들의 값을 출력하라

const (
	a     = 42
	b int = 44
)

func main() {
	fmt.Println(a, b)
}
