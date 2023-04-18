package dog

import (
	"fmt"
	"github.com/praslar/meraki/meraki/functions/animal"
)

type Dog struct {
	animal.Animal // embedded type
	BarkStrength int
}


func (dog *Dog) MakeNoise() {
	barkStrength := dog.BarkStrength
	//
	//if dog.Animal. == true {
	//	barkStrength = barkStrength * 5
	//}

	for bark := 0; bark < barkStrength; bark++ {
		fmt.Printf("BARK ")
	}

	fmt.Println("...")
}
