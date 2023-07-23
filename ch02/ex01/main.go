package main

import "fmt"

// 값을 따로 지정하지 않으면 false
var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 7
	b := 42
	fmt.Println(a == b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

}
