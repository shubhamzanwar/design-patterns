package main

type iPizza interface {
	GetPrice() float64
	GetName() string
	GetToppings() []string
}

type pizza struct {
	name     string
	price    float64
	toppings []string
}

func (p *pizza) GetName() string {
	return p.name
}

func (p *pizza) GetPrice() float64 {
	return p.price
}

func (p *pizza) GetToppings() []string {
	return p.toppings
}
