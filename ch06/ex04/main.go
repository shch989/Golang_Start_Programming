package main

import "fmt"

// defer는 함수가 다 실행되고 실행되도록 해준다.

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
