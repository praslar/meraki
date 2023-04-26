package utils

import (
	"bufio"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"meraki/configs"
	"os"
	"strconv"
	"strings"
)

func ConnectDB(conf configs.ConfigDB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", conf.Host, conf.Username, conf.Password, conf.Dbname, conf.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InMenuChinh() {
	fmt.Println("===================He Thong Chuc Nang=======================")
	fmt.Println("0. Migrate database")
	fmt.Println("1. Đăng ký bằng email.")
	fmt.Println("2. Đăng nhập bằng email.")
	fmt.Println("3. Save.")
	fmt.Println("============================================================")
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
