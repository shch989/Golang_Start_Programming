package main

import "fmt"

// person 구조체 정의
type person struct {
	first string
	last  string
}

// secretAgent 구조체 정의, person 구조체를 포함
type secretAgent struct {
	person // person 구조체를 포함하여 필드 재사용 (상속과 비슷한 효과)
	ltk    bool
}

// person 구조체의 speak() 메서드 구현
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

// secretAgent 구조체의 speak() 메서드 구현
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- the secretAgent speak")
}

// human 인터페이스 정의 (speak() 메서드만 가지도록)
type human interface {
	speak()
}

// bar 함수 정의, human 인터페이스를 매개변수로 받음
func bar(h human) {
	// 인자로 전달된 객체의 유형에 따라 다른 동작 수행
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

// hotdog 타입 정의 (사용자 정의 타입)
type hotdog int

func main() {
	// secretAgent 구조체 인스턴스 생성
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	// secretAgent 구조체 인스턴스 생성
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	// person 구조체 인스턴스 생성
	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	// speak() 메서드 호출
	sa1.speak() // 결과: "I am James Bond - the secretAgent speak"
	sa2.speak() // 결과: "I am Miss Moneypenny - the secretAgent speak"

	fmt.Println(p1) // 결과: {Dr. Yes}

	// bar 함수 호출하여 다양한 객체의 speak() 메서드 출력
	bar(sa1) // 결과: "I was passed into bar James"
	bar(sa2) // 결과: "I was passed into bar Miss"
	bar(p1)  // 결과: "I was passed into bar Dr."

	// hotdog 타입 변수 생성 및 타입 변환
	var x hotdog = 42
	fmt.Println(x)        // 결과: 42
	fmt.Printf("%T\n", x) // 결과: main.hotdog
	var y int
	y = int(x)
	fmt.Println(y)        // 결과: 42
	fmt.Printf("%T\n", y) // 결과: int
}
