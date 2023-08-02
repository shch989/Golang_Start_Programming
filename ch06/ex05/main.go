package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sal1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sal2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	fmt.Println(sal1)
	fmt.Println(sal2)

	sal1.speak()
	sal2.speak()
}
