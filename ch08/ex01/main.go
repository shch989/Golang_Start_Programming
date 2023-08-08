package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	// people 변수에 저장된 데이터를 JSON 형식으로 변환 후 bs 변수에 저장
	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
