package main

import "fmt"

var y = 42
var z = "Shaken, not stirred"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	// Type Error!
	// var z에서 문자값을 변수로 지정하였기 때문에 String 타입만 변수지정 가능
	// z = 43
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)
}
