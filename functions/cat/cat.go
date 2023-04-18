package cat

import (
	"fmt"
	"github.com/praslar/meraki/meraki/functions/animal"
)

type Cat struct {
	animal.Animal // embedded
	MeowStrength int
}



func (cat *Cat) MakeNoise() {
	meowStrength := cat.MeowStrength


	for meow := 0; meow < meowStrength; meow++ {
		fmt.Printf("MEOW ")
	}

	fmt.Println("...")
}