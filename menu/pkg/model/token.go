package model

type Token struct {
	Address  string  `json:"address" gorm:"primaryKey"`
	Symbol   string  `json:"symbol"`
	Balances float64 `json:"balances"`
}
