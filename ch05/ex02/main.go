package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sal := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sal)
	fmt.Println(p2)

	fmt.Println(sal.first, sal.last, sal.age, sal.ltk)
	fmt.Println(p2.first, p2.last, p2.age)
}
