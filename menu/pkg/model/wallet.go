package model

type Wallet struct {
	Address string `json:"address" gorm:"primaryKey"`
	Name    string `json:"name"`
}
