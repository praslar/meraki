package main

import "fmt"

func main() {

	// i= 0
	dayso := []int{1, 9, 8, 6, 3, 99, 5, 1, 3}
	dayso := []int{1, 8, 9, 6, 3, 99, 5, 1, 3}
	dayso := []int{1, 8, 6, 9, 3, 99, 5, 1, 3}
	dayso := []int{1, 8, 6, 3, 9, 99, 5, 1, 3}
	dayso := []int{1, 8, 6, 3, 9, 5, 99, 1, 3}
	dayso := []int{1, 8, 6, 3, 9, 5, 1, 99, 3}

	// i = 1
	dayso := []int{1, 6, 3, 8, 5, 1, 9, 3, 99}

	// i = 2
	dayso := []int{1, 3, 6, 5, 1, 8, 3, 9, 99}
	// i = 4
	dayso := []int{1, 1, 3, 5, 6, 8, 3, 9, 99}

	i = 0

	j = 0
	j = 1
	j = 2
	j = 3
	j = 4
	j = 5
	j = 6
	j = 7

	a, b := 1, 2
	fmt.Println(a, b)
	//len = 9
	fmt.Println("Sap xep theo thu tu tang dan")

	// bubble sort
	for i := 0; i < len(dayso)-1; i++ {
		for j := 0; j < len(dayso)-i-1; j++ {

			if dayso[j] > dayso[j+1] {

				tmp := dayso[j]
				dayso[j] = dayso[j+1]
				dayso[j+1] = tmp
			}
		}
	}

	fmt.Println(dayso)
}
