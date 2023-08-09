package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

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
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Hello World",
			"Hello World",
			"Hello World",
			"Hello World",
		},
	}

	u3 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Hello World",
			"Hello World",
			"Hello World",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error: ", err)
	}
}
