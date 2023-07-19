package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	// 타입
	fmt.Printf("%T\n", y)
	// 2진수
	fmt.Printf("%b\n", y)
	// 16진수
	fmt.Printf("%x\n", y)
	// 0x 16진수
	fmt.Printf("%#x\n", y)

	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
	// 10진수
	fmt.Printf("%v", y)
}
