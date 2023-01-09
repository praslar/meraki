package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	XuatMenu()
	for {
		fmt.Print("Nhap chuc nang can thuc hien: ")
		option := NhapOption()
		if option == 9 {
			fmt.Println(" Đang thoát chương trình...")
			break
		}
		XuOption(option)
	}

	fmt.Println("Bye")
}

func NhapOption() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	// handle
	if err != nil {
		fmt.Println("Đã xảy ra lỗi: ", err)
		return -1
	}

	// convert string to int
	option, err := strconv.Atoi(input[:len(input)-2])
	if err != nil {
		fmt.Println("Đã xảy ra lỗi: ", err)
		return -1
	}
	return option
}

func XuatMenu() {
	fmt.Println("===================He Thong Chuc Nang=======================")
	fmt.Println("0. Tạo wallet")
	fmt.Println("1. Xóa wallet")
	fmt.Println("2. Hiển thị lại danh sách hệ thống chức năng")
	fmt.Println("9. Để thoát chương trình")
	fmt.Println("============================================================")
	fmt.Println("Vui lòng chọn option từ 1 - 9")
}

func XuOption(option int) {
	switch option {
	case 0:
		fmt.Println("Bạn đã chọn options: ", option)
		fmt.Println("Chương trình đang thực hiện chức năng ...")
		fmt.Println("Done")
		fmt.Println("------------------------------------------")
	case 1:
		fmt.Println("Bạn đã chọn options: ", option)
		fmt.Println("Chương trình đang thực hiện tải về ")
		fmt.Println("Done")
		fmt.Println("------------------------------------------")
	case 2:
		XuatMenu()
	default:
		fmt.Println("Option bạn chọn không có trong menu. Vui lòng chọn lại")
		fmt.Println("------------------------------------------------------")
		break
	}
}
