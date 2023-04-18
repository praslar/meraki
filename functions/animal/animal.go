package animal

type Animal struct {
	Name string
	mean bool
}

func (animal *Animal) GetName() string {
	return animal.Name
}

func (animal *Animal) SetMean(foodType string) {
	if foodType == "meat" {
		animal.mean = true
	}
	if foodType == "weed" {
		animal.mean = false
	}
}