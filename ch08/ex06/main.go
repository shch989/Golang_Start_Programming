package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 원래 비밀번호
	s := "password123"
	// 비밀번호 해시 생성
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(string(bs))

	// 로그인 시도할 비밀번호
	loginPword1 := `password123`

	// 비밀번호 해시와 로그인 시도한 비밀번호 비교
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("로그인할 수 없습니다")
		return
	}
	fmt.Println("로그인되었습니다")
}
