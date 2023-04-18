package main

import (
	"fmt"
	"meraki/configs"
	"meraki/pkg/model"
	"meraki/pkg/utils"
)

func main() {
	fmt.Println("Hello, welcome to meraki wallet project")

	// connect to db
	pg, err := utils.ConnectDB(configs.ConfigDB{
		Host:     "172.104.41.46",
		Port:     "5432",
		Username: "nmadmin",
		Password: "L9nLshJ3F7NqJgu7",
		Dbname:   "meraki",
	})
	// error handling
	if err != nil {
		fmt.Println("Đã có lỗi xảy ra: ", err)
		return
	}

	if err := pg.AutoMigrate(
		&model.User{},
		&model.UserHasWallet{},
		&model.Wallet{},
		&model.Transaction{},
		&model.Token{},
	); err != nil {
		fmt.Println("Đã có lỗi xảy ra: ", err)
		return
	}

	fmt.Println("Da migrate thanh cong")
}
