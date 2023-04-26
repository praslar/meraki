package handler

import (
	"fmt"
	"gorm.io/gorm"
	"meraki/pkg/model"
	"meraki/pkg/utils"
	"net/mail"
)

type UserHandler struct {
	DbConnection *gorm.DB
}

func NewUserHandler(db *gorm.DB) UserHandler {
	return UserHandler{
		DbConnection: db,
	}
}

func (h *UserHandler) Register() error {

	fmt.Print("Mời bạn nhập email: ")
	newEmail := utils.GetInputString()
	if !validEmail(newEmail) {
		return fmt.Errorf("lỗi định dạnh email")
	}

	fmt.Print("Mời bạn nhập password: ")
	password := utils.GetInputString()

	if h.CheckEmailExist(newEmail) {
		return fmt.Errorf("email đã tồn tại")
	}

	if err := h.DbConnection.Create(&model.User{
		Email:    newEmail,
		Password: password,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (h *UserHandler) Login() {
	//h.DbConnection.Find()
}

func (h *UserHandler) Logout() {

}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (h *UserHandler) CheckEmailExist(newEmail string) bool {
	rs := model.User{}
	err := h.DbConnection.Model(&model.User{}).Where("email = ?", newEmail).First(&rs).Error
	if err != nil {
		// email chưa tôn tại, nên không tìm tim thấy
		return false
	}
	return true
}
