package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/mail"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	Email   string   `json:"email"`
	Wallets []Wallet `json:"wallets"`
}

type Wallet struct {
	Address string  `json:"address"`
	Name    string  `json:"name"`
	Tokens  []Token `json:"tokens"`
}

type Token struct {
	Symbol   string  `json:"symbol"`
	Balances float64 `json:"balances"`
}

func InMenuChinh() {
	fmt.Println("===================He Thong Chuc Nang=======================")
	fmt.Println("0. Xem list user hiện tại")
	fmt.Println("1. Đăng ký bằng email.")
	fmt.Println("2. Đăng nhập bằng email.")
	fmt.Println("3. Save.")

	fmt.Println("============================================================")
}

func InMenuAddToken() {
	fmt.Println("=================== Tạo Token Cho wallet =======================")
	fmt.Println("0. Add Token Mới")
	fmt.Println("1. Về menu chính.")
	fmt.Println("============================================================")
}

func InMenuWallet() {
	fmt.Println("=================== Tạo Token Cho wallet =======================")
	fmt.Println("1. Tạo thêm wallet")
	fmt.Println("2. Sửa name wallet")
	fmt.Println("3. Xóa wallet theo address")
	fmt.Println("0. Quay về màn hình login")
	fmt.Println("============================================================")
}

func main() {
	// TODO : listUsers = DocListTuFile()

	var listUsers []User

	data, err := read()
	if err != nil {
		fmt.Println("Read file error!")
		listUsers = nil
	}

	listUsers = data

	for {
		InMenuChinh()
		fmt.Print("Nhap chuc nang can thuc hien: ")
		option := NhapOption()
		switch option {
		case 0:
			// clear console
			fmt.Print("\033[H\033[2J")
			fmt.Println("List user hiện tại: ")
			PrintThongTin(listUsers)
		case 1: // Đăng ký email
			fmt.Print("Mời bạn nhập email: ")
			newEmail := GetInputString()
			if !validEmail(newEmail) {
				fmt.Println("Lỗi định dạnh email")
				break
			}
			if !checkEmailExist(newEmail, listUsers) {
				fmt.Println("Email đã tồn tại")
				break
			}
			var initTokens []Token
			optionMenuPhu := 0
			for {
				InMenuAddToken()
				fmt.Print("Chọn option: ")
				optionMenuPhu = NhapOption()

				if optionMenuPhu == 0 {
					fmt.Print("Nhập symbol: ")
					symbol := GetInputString()
					fmt.Print("Nhập balance: ")
					balance := GetInputNumber()
					initTokens = append(initTokens, Token{
						Symbol:   symbol,
						Balances: balance,
					})
				}

				if optionMenuPhu == 1 {
					break
				}
			}

			// init wallet
			initWallet := Wallet{
				Address: uuid.NewString(),
				Tokens:  initTokens,
			}

			listUsers = append(listUsers, User{
				Email:   newEmail,
				Wallets: []Wallet{initWallet},
			})

			PrintThongTin(listUsers)
		case 2:
			// TODO:
			// kiểm tra email có tồn tai khong
			// nếu email có thì hiển thị menu thao tac voi wallet
			fmt.Print("Mời bạn nhập email: ")
			currentEmail := GetInputString()
			if !validEmail(currentEmail) {
				fmt.Println("Lỗi định dạnh email")
				break
			}
			isLogin := false
			currentUserIndex := -1
			for index, user := range listUsers {
				if currentEmail == user.Email {
					fmt.Println("Bạn đã đăng nhập thành công")
					currentUserIndex = index
					fmt.Println("Index của ban là: ", currentUserIndex)
					isLogin = true
					break
				}
			}

			if isLogin == false {
				fmt.Println("Đăng nhập thất bại")
				break
			}

			fmt.Println("Làm gì làm tiếp đi")

			fmt.Printf("Thong tin ví của %s là: ", currentEmail)
			PrintThongTinUser(listUsers[currentUserIndex])
			for {
				InMenuWallet()
				optionMenuWallet := GetInputNumber()
				if optionMenuWallet == 0 {
					fmt.Println("Logout")
					break
				}
				if optionMenuWallet == 1 {
					// thêm wallet
					// khi thêm wallet thì mặc định là không có token nào,
					// nếu muốn thêm token thì chọn option add Token (chưa hỗ trợ)
					newWallet := Wallet{
						Address: uuid.NewString(),
						Tokens:  []Token{},
					}
					listUsers[currentUserIndex].Wallets = append(listUsers[currentUserIndex].Wallets, newWallet)
					fmt.Print("\033[H\033[2J")
					fmt.Println("bạn đã tạo ví mới thành công!")
					PrintThongTinUser(listUsers[currentUserIndex])
				}
				if optionMenuWallet == 2 {
					// sửa name wallet
					fmt.Println("Bạn muốn đổi email thành gì: ")
					newEmail := GetInputString()
					if !validEmail(newEmail) {
						break
					}
					listUsers[currentUserIndex].Email = newEmail
					fmt.Print("\033[H\033[2J")
					fmt.Println("bạn đã đổi email mới thành: ", newEmail)
					PrintThongTinUser(listUsers[currentUserIndex])
				}
				if optionMenuWallet == 3 {
					// xóa wallet
					fmt.Print("Bạn muốn xóa address nào: ")
					deleteAddress := GetInputString()
					isDeleted := false
					for i, wallet := range listUsers[currentUserIndex].Wallets {
						if deleteAddress == wallet.Address {
							listUsers[currentUserIndex].Wallets = append(listUsers[currentUserIndex].Wallets[:i], listUsers[currentUserIndex].Wallets[i+1:]...)
							isDeleted = true
							break
						}
					}
					if !isDeleted {
						fmt.Println("Không tìm thấy địa chỉ để xóa")
						break
					}

					fmt.Print("\033[H\033[2J")
					fmt.Println("bạn đã xóa wallet: ", deleteAddress)
					PrintThongTinUser(listUsers[currentUserIndex])
				}
			}
		case 3:
			fmt.Println(" Saving....")
			fmt.Println("Please do not close the app during saving...")
			err := save(listUsers)
			if err != nil {
				fmt.Println(" ERROR: cannot save data: ", err)
			}
			fmt.Println("Saved.")
		default:
			fmt.Println("Option bạn chọn không có trong menu. Vui lòng chọn lại")
			fmt.Println("------------------------------------------------------")
		}

	}
}

func GetInputString() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Đã xảy ra lỗi: ", err)
		return ""
	}
	inputRemovedEnter := strings.Trim(input, "\r\n")
	return inputRemovedEnter
}

func GetInputNumber() float64 {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Đã xảy ra lỗi: ", err)
		return 0
	}
	inputRemovedEnter := strings.Trim(input, "\r\n")
	toFloatBalance, err := strconv.ParseFloat(inputRemovedEnter, 64)
	if err != nil {
		return 0
	}
	return toFloatBalance
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

func read() ([]User, error) {

	var listUser []User
	dat, err := os.ReadFile("./data.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dat, &listUser)
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func save(data []User) error {
	tmp, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile("./data.json", tmp, 0644)
	if err != nil {
		return err
	}

	return nil
}

func PrintThongTin(listUser []User) {
	for _, user := range listUser {
		fmt.Println("Email: ", user.Email)
		for _, wallet := range user.Wallets {
			fmt.Println("Address: ", wallet.Address)
			for _, token := range wallet.Tokens {
				fmt.Println("Symbol: ", token.Symbol)
				fmt.Println("Balance: ", token.Balances)
			}
		}
	}
}
func PrintThongTinUser(user User) {
	fmt.Println("Email: ", user.Email)
	for _, wallet := range user.Wallets {
		fmt.Println("Address: ", wallet.Address)
		for _, token := range wallet.Tokens {
			fmt.Println("Symbol: ", token.Symbol)
			fmt.Println("Balance: ", token.Balances)
		}
	}
}
