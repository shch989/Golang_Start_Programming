package main

import "fmt"

// 슬라이스를 만들어서 한국의 모든 행정구역의 이름을 저장하시오(make와 append를 사용)
// range 절을 사용하지 않고 이 값들을 인덱스 위치와 함께 출력하시오

func main() {
	x := make([]string, 50, 50)
	x = []string{`서울특별시`, `부산광역시`, `대구광역시`, `인천광역시`, `광주광역시`, `대전광역시`, `울산광역시`, `세종특별자치시`, `경기도`, `강원특별자치도`, `충청북도`, `충청남도`, `전라북도`, `전라남도`, `경상북도`, `경상남도`, `제주특별자치도`}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
