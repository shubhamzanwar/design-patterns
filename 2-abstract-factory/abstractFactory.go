package main

import "fmt"

type iPizzaFactory interface {
	createPizza() iPizza
	createGarlicBread() iGarlicBread
}

func getFactory(chain string) (iPizzaFactory, error) {
	if chain == "P" {
		return &pizzaHutFactory{}, nil
	}
	if chain == "D" {
		return &dominosFactory{}, nil
	}
	return nil, fmt.Errorf("Enter a valid chain type next time")
}
