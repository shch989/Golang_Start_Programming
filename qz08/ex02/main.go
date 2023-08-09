package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	s := `[{"First": "James", "Last": "Bond", "Age": 32, "Sayings": ["Hello", "Byebye"]}, {"First": "James", "Last": "Bond", "Age": 32, "Sayings": ["Hello", "Byebye"]}, {"First": "James", "Last": "Bond", "Age": 32, "Sayings": ["Hello", "Byebye"]}]`
	fmt.Println(s)

	var people []person

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, person := range people {
		fmt.Println("Person #", i)
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}
