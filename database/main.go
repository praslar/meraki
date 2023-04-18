package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello")
}

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Wallet struct {
	Address string    `json:"address"`
	UserId  uuid.UUID `json:"user_id"`
}

type Token struct {
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
}

type Transaction struct {
	From    string  `json:"from"`
	To      string  `json:"to"`
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}
