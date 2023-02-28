package main

import (
	"encoding/json"
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

func main() {
	listUser := []User{
		{
			Email: "thang",
			Wallets: []Wallet{
				{
					Address: "0xabc",
					Tokens: []Token{
						{"eth", 100},
						{"btc", 10},
					},
				},
			},
		},
		{
			Email: "thinh",
			Wallets: []Wallet{
				{
					Address: "0xzxy",
					Tokens: []Token{
						{"ada", 1000},
						{"dogecoin", 1000000},
						{"doge", 4634534},
						{"zyn", 1231.123123},
						{"dogecoin", 9999},
					},
				},
			},
		},
	}

	tmp, err := json.Marshal(listUser)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./data.json", tmp, 0644)
	if err != nil {
		panic(err)
	}

}
