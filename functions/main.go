package main

import (
	"fmt"
	"github.com/praslar/meraki/meraki/functions/animal"
	"github.com/praslar/meraki/meraki/functions/cat"
)


func main()  {
	sua := &cat.Cat{
		Animal:       animal.Animal{
			Name: "Sữa",
		},
		MeowStrength: 10,
	}



	fmt.Printf("Thông tin cat: %v, animal: %v \n", sua.MeowStrength, sua.Animal)
	// =======================================================================================


	sua.SetMean("meat")
	fmt.Printf("Thông tin cat: %v, animal: %v \n", sua.MeowStrength, sua.Animal)
}

