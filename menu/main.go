package main

import (
	"bufio"
	"fmt"
	"net/mail"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Email   string   `json:"email"`
	Wallets []Wallet `json:"wallets"`
}

type Wallet struct {
	Address string  `json:"address"`
	Tokens  []Token `json:"tokens"`
}

type Token struct {
	Symbol   string  `json:"symbol"`
	Balances float64 `json:"balances"`
}

func main() {
	// TODO : listUsers = DocListTuFile()
	fmt.Println("===================He Thong Chuc Nang=======================")
	fmt.Println("0. Xem list user hiện tại")
	fmt.Println("1. Đăng ký bằng email.")
	fmt.Println("============================================================")

	var listUsers []User
	for {
		fmt.Print("Nhap chuc nang can thuc hien: ")
		option := NhapOption()
		switch option {
		case 0:
			fmt.Println("List user hiện tại: ")
			fmt.Println(listUsers)
		case 1:
			fmt.Print("Mời bạn nhập email: ")
			newEmail := GetInput()
			if !validEmail(newEmail) {
				fmt.Println("Lỗi định dạnh email")
				break
			}
			if !checkEmailExist(newEmail, listUsers) {
				fmt.Println("Email đã tồn tại")
				break
			}
			listUsers = append(listUsers, User{
				Email:   newEmail,
				Wallets: nil,
			})
			fmt.Println(listUsers)
		default:
			fmt.Println("Option bạn chọn không có trong menu. Vui lòng chọn lại")
			fmt.Println("------------------------------------------------------")
		}

	}
}

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Đã xảy ra lỗi: ", err)
		return ""
	}
	inputRemovedEnter := strings.Trim(input, "\r\n")
	return inputRemovedEnter
}

func NhapOption() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Đã xảy ra lỗi: ", err)
		return -1
	}

	inputRemovedEnter := strings.Trim(input, "\r\n")
	option, err := strconv.Atoi(inputRemovedEnter)
	if err != nil {
		fmt.Println("Đã xảy ra lỗi: ", err)
		return -1
	}
	return option
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func checkEmailExist(newEmail string, listUsers []User) bool {
	for _, user := range listUsers {
		if newEmail == user.Email {
			return false
		}
	}
	return true
}
