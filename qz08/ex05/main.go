package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Hello World",
			"Hello World",
		},
	}

	u2 := user{
		First: "John",
		Last:  "Bob",
		Age:   42,
		Sayings: []string{
			"Hello World",
			"Hello World",
			"Hello World",
			"Hello World",
		},
	}

	u3 := user{
		First: "Michael",
		Last:  "Jenny",
		Age:   26,
		Sayings: []string{
			"Hello World",
			"Hello World",
			"Hello World",
		},
	}

	users := []user{u1, u2, u3}

	sort.Sort(ByAge(users))

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
		fmt.Println("\t")
	}

	fmt.Println("------------------------------------------")

	sort.Sort(ByLast(users))

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
		fmt.Println("\t")
	}

}
