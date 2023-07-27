package main

import "fmt"

// switch문을 사용하여 문자열형의 변수 favSport를 사용하여 식을 만드시오

func main() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to the hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}
