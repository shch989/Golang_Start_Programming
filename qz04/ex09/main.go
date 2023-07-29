package main

import "fmt"

// 이전 실습에서 맵에 새로운 항목을 추가한다.
// range를 이용해 맵을 출력한다.

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`fleming_lan`] = []string{`steaks`, `clgars`, `espionage`}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
