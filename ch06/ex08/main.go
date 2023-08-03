package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")

	f := func(x int) {
		fmt.Println("the year big brother started watching:", x)
	}

	f(1984)
}
