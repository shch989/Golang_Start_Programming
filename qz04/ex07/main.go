package main

import "fmt"

// [][]string 형태의 문자열 슬라이스를 만드시오
// "James", "Bond", "Shaken, not Stirred"
// "Miss", "Moneypenny", "Hellooooo, James."
// 위의 두 데이터를 다차원 슬라이스에 저장하여 출력하시오

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not Stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Hellooooo, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, v := range xs {
			fmt.Printf("\t index: %v\t value: %v\n", j, v)
		}
	}
}
