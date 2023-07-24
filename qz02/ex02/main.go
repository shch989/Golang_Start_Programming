package main

import "fmt"

// 다음과 같은 연산자를 이용하고 표현식을 써서 값을 할당하시오
// g. ==
// h. <=
// i. >=
// j. !=
// k. <
// l. >

func main() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)

	fmt.Println(a, b, c, d, e, f)
}
