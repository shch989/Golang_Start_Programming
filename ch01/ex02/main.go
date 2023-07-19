package main

import (
	"fmt"
)

func main() {
	n, e := fmt.Println("Hello, playground", 42, true)
	// 26bit 출력
	fmt.Println(n)
	// err nil 출력
	fmt.Println(e)
}
