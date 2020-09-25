package main

type iGarlicBread interface {
	GetPrice() float64
	GetName() string
}

type garlicBread struct {
	name  string
	price float64
}

func (g *garlicBread) GetName() string {
	return g.name
}

func (g *garlicBread) GetPrice() float64 {
	return g.price
}
