package main

import "fmt"

// iota : 연속된 정수값을 생성

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
