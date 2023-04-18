package main

import "fmt"

type Vu struct {
	Name string
	Age int
}

func main() {

	a := Vu{
		Name: "Vu",
	}

	b := &a
	b.Name = "thang"

	fmt.Println(a)
	fmt.Println(b)
}


