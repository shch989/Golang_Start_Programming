package main

import "fmt"

// vehicle 타입을 생성하여 문의 갯수와 색상을 나타내는 타입을 설정하시오
// vehicle의 타입을 가지고 4륜 바퀴인지 아닌지를 나타내는 truck 타입을 생성하시오
// vehicle의 타입을 가지고 세련된 차량인지 아닌지를 나타내는 sedan 타입을 생성하시오
// truck과 sedan 타입을 사용하여 데이터를 추가하고 출력하시오

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}
	fmt.Println(t)

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(s)
}
