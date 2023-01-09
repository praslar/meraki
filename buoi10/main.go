package main

import "fmt"

func main() {
	// y := x/3 + 15
	for x := 1000; x < 1010; x++ {
		y := x/3 + 15
		fmt.Print("(", x, ";", y, ")", "\t")
	}
}
