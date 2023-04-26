package main

import (
	"fmt"
	"meraki/configs"
	"meraki/pkg/handler"
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

	// khởi tạo user handler
	userHandler := handler.NewUserHandler(pg)

	for {

		// Lấy request của user
		utils.InMenuChinh()
		fmt.Print("Nhap chuc nang can thuc hien: ")
		option := utils.NhapOption()

		// xử lý logic
		switch option {
		case 0:
			// tạo bảng
			_ = pg.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
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
		case 1:
			if err := userHandler.Register(); err != nil {
				fmt.Println("Lỗi tạo user: ", err)
			} else {
				fmt.Println("Tạo user thành công")
			}
		}
	}

}
