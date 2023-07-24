package main

import "fmt"

// x의 값을 10진수 2진수 16진수로 출력하시오
// x = 42

func main() {
	x := 42
	fmt.Printf("%d\t%b\t%#x", x, x, x)
}
