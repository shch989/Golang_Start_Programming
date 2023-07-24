package main

import "fmt"

// 원시 문자열 리터럴 변수를 만들고 출력하시오

func main() {
	a := `here is something
	as
	a
	raw string
	literal
	"you see"
	another thing`
	fmt.Println(a)
}
