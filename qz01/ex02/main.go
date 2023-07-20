package main

import "fmt"

// 다음의 식별자를 변수에 사용하고 변수들이 다음과 같은 타입이 되도록 하시오.
// int 타입의 x
// string 타입의 y
// bool 타입의 z

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
