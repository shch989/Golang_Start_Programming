package main

import "fmt"

// 문자열 키를 성_이름의 형태로 맵을 만들고
// 값은 문자열 슬라이스 타입으로 한다.
// 그들이 좋아하는 것을 저장하여 맵에 저장한다.
// 해당 값을 슬라이스의 인덱스와 함께 출력한다.

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
