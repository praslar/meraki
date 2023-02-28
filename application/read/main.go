package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Email   string   `json:"email"`
	Wallets []Wallet `json:"wallets"`
}

type Wallet struct {
	Address string  `json:"address"`
	Tokens  []Token `json:"tokens"`
}

type Token struct {
	Symbol   string  `json:"symbol"`
	Balances float64 `json:"balances"`
}

// struct -> json encoded byte -> write file

// read file -> json encoded byte  -> parse struct
func main() {
	fmt.Println("Hello")

	listUser := []User{}

	dat, err := os.ReadFile("./data.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(dat, &listUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(listUser)
	for _, testStruct := range listUser {
		fmt.Println(testStruct)
	}
}
