package main

import "fmt"

func main() {
	n := "Bond"
	switch n {
	case "Moneypenny", "bond", "Dr No":
		fmt.Println("miss money or bond or dr no")
	case "Bond":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}

}
