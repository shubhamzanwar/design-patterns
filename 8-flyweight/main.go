package main

import "fmt"

const (
	// NikeAirForce is the identifier for the air force sneaker
	NikeAirForce = "Nike Air Force 1"
	// AdidasSuperstar is the identifier for the Superstar sneaker
	AdidasSuperstar = "Adidas Superstar"
)

func main() {
	sneaker1 := Sneaker{
		shoeType: NikeAirForce,
		size:     9,
		color:    "white",
	}

	sneaker2 := Sneaker{
		shoeType: AdidasSuperstar,
		size:     8,
		color:    "black",
	}

	fmt.Println(sneaker1.describe())
	fmt.Println(sneaker2.describe())
}
