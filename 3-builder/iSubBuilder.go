package main

type iSubBuilder interface {
	setBread()
	setCheese()
	setToppings()
	setSauces()
	getSub() sub
}
