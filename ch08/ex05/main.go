package main

import (
	"fmt"
	"sort"
)

// Person 구조체는 사람의 이름과 나이를 저장합니다.
type Person struct {
	Name string
	Age  int
}

// String 메서드는 Person 구조체를 문자열로 표현합니다.
func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByName은 이름에 따라 정렬하기 위한 타입입니다.
type ByName []Person

// Len 메서드는 ByName 슬라이스의 길이를 반환합니다.
func (bn ByName) Len() int {
	return len(bn)
}

// Swap 메서드는 ByName 슬라이스 내의 두 요소를 교환합니다.
func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

// Less 메서드는 ByName 슬라이스 내의 두 요소를 비교하여 정렬 순서를 결정합니다.
func (bn ByName) Less(i, j int) bool {
	return bn[i].Name < bn[j].Name
}

func main() {
	// 사람들의 정보를 생성합니다.
	p1 := Person{"Bob", 31}
	p2 := Person{"John", 42}
	p3 := Person{"Michael", 17}
	p4 := Person{"Jenny", 26}

	// Person 구조체를 슬라이스로 묶어줍니다.
	people := []Person{p1, p2, p3, p4}

	// 초기 사람들의 정보 출력
	fmt.Println(people)

	// 이름에 따라 정렬
	sort.Sort(ByName(people))

	// 정렬된 사람들의 정보 출력
	fmt.Println(people)
}
