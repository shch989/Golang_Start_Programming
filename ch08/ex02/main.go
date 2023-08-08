package main

import (
	"encoding/json"
	"fmt"
)

// person 구조체 정의
type person struct {
	First string `json:"First"` // JSON 필드 이름이 First인 필드에 대응
	Last  string `json:"Last"`  // JSON 필드 이름이 Last인 필드에 대응
	Age   int    `json:"Age"`   // JSON 필드 이름이 Age인 필드에 대응
}

func main() {
	// JSON 문자열
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)

	// 문자열 및 바이트 슬라이스의 타입 출력
	fmt.Printf("%T\n", s)  // string
	fmt.Printf("%T\n", bs) // []uint8 (바이트 슬라이스는 uint8의 별칭)

	// person 슬라이스 정의
	people := []person{}

	// JSON 파싱
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)

	// 파싱된 데이터 출력
	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
