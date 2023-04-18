package model

type Transaction struct {
	ID string `json:"id" gorm:"primaryKey"`

	FromAddress string `json:"from_address"`
	WalletFrom  Wallet `json:"wallet_from" gorm:"foreignKey:from_address;references:address"`

	ToAddress string `json:"to_address"`
	WalletTo  Wallet `json:"wallet_to" gorm:"foreignKey:to_address;references:address"`

	TokenAddress string  `json:"token_address"`
	Token        Token   `json:"token" gorm:"foreignKey:token_address;references:address"`
	Amount       float64 `json:"amount"`
}
