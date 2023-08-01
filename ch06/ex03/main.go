package main

import "fmt"

func main() {
	x, y := sum("james", 4, 5, 6)
	fmt.Println(x, y)
}

func sum(s string, x ...int) (string, int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	name := s
	for _, v := range x {
		sum += v
	}

	return name, sum
}
