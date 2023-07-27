package main

import "fmt"

// switch 문에서 식을 작성하지 않고 프로그램을 작성하시오

func main() {
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}
