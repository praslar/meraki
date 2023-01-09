package main

import (
	"fmt"
)

func main() {
	fmt.Println("------------------------------------------------------------")
	fmt.Println("-----------------------WELCOME TO--------------------------")
	fmt.Println("-----------------MERAKI BLOCKCHAIN--------------------------")
	fmt.Println("1. Enter any letter to print to the console")
	fmt.Println("2. Press q to exit")

	fmt.Println("---------------this blockchain created by meraki team--------")

	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)
}
