package main

import "fmt"

// 이전 예제의 코드에서 다음의 값을 각각의 변수에 지정하여라
// x. 42
// y. "James Bond"
// z. true
// fmt.Sprintf를 이용해서 모든 값을 하나의 문자열로 프린트하여라

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
