package main

import "fmt"

// iota를 이용해 최근 4년에 대한 4개의 상수를 만들어라
// 해당 상수들을 출력하라

// const (
// 	a = 2020 + iota
// 	b = 2020 + iota
// 	c = 2020 + iota
// 	d = 2020 + iota
// )

const (
	a = 2020 + iota
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
