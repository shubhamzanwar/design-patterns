package main

import "fmt"

// Sneaker is the struct to hold sneaker specific details
type Sneaker struct {
	shoeType string
	size     int
	color    string
}

func (s *Sneaker) describe() string {
	details := getShoeDetails(s.shoeType)
	return fmt.Sprintf("The %s is of size %d, color %s and costs %.2f. The Manufacturer is %s", details.name, s.size, s.color, details.price, details.brand)
}
