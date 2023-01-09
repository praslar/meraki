package main

import (
	"fmt"
)

func main() {

	// nếu string gặp chũ T thì là nút tab trên bàn phím

	// abcTzxc --> abc	zxc
	result := Solution("abcTzxc")

	fmt.Println(result)
}

func Solution(S string) string {
	var ketQua []rune
	for _, v := range S {
		ketQua = append(ketQua, v)
	}
	return string(ketQua)
}
