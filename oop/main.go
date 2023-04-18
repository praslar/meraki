package main

import (
	"fmt"
	"wallet/user"
)


func main() {
	///
	thangPham := user.User{
		Email:    "pxthang",
		Password: "123456",
		Birthday: "11/08/1997",
		InitBalance: 1000000,
	}
	//
	//
	//err := thangPham.CreateUser(user.User{
	//	Email:    "vu",
	//	Password: "123456",
	//	Birthday: "1/1/1997",
	//	InitBalance: 1000000,
	//})

	newUser := thangPham.CreateUser("vu@gmail.com","1/1/1997")
	newUser2 := thangPham.CreateUser("thinh@gmail.com","1/1/1997")
	newUser3 := thangPham.CreateUser("lo@gmail.com","1/1/1997")
	fmt.Println(newUser,newUser2,newUser3)


}


