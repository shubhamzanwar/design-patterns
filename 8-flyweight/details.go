package main

type sneakerDetails struct {
	name        string
	brand       string
	price       float64
	materials   []string
	description string
	image       string
	laced       bool
}

// NikeAirForceDetails is the details for air force sneaker
var NikeAirForceDetails = sneakerDetails{
	name:        "Nike Air Force 1",
	brand:       "Nike",
	price:       144.99,
	materials:   []string{"leather", "rubber"},
	description: "Originally released in 1982, the Nike Air Force 1 was the first Nike model to feature Air technology.",
	image:       "https://via.placeholder.com/50",
	laced:       true,
}

// AdidasSuperstarDetails is the details for Superstar sneaker
var AdidasSuperstarDetails = sneakerDetails{
	name:        "Adidas Superstar",
	brand:       "Adidas",
	price:       85.99,
	materials:   []string{"leather", "rubber"},
	description: "Originally made for basketball courts in the '70s.",
	image:       "https://via.placeholder.com/50",
	laced:       true,
}

// DetailsMap stores the details against the sneaker type
var DetailsMap = map[string]sneakerDetails{
	NikeAirForce:    NikeAirForceDetails,
	AdidasSuperstar: AdidasSuperstarDetails,
}

func getShoeDetails(shoeType string) sneakerDetails {
	return DetailsMap[shoeType]
}
