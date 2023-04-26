package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type UserHasWallet struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`

	UserID uuid.UUID `json:"user_id"`
	User   User      `json:"user" gorm:"foreignKey:user_id;;references:id"`

	WalletAddress string `json:"wallet_address"`
	Wallet        Wallet `json:"wallet" gorm:"foreignKey:wallet_address;;references:address"`
}

// TODO: USER
// 1. Register
// 2. Login
