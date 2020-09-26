package main

type veggieDelightBuilder struct {
	sub
}

func (v *veggieDelightBuilder) setBread() {
	v.sub.bread = "parmesan oregano"
}

func (v *veggieDelightBuilder) setCheese() {
	v.sub.hasCheese = false
}

func (v *veggieDelightBuilder) setToppings() {
	v.sub.toppings = []string{"olives", "tomatoes", "onions", "jalapeños"}
}

func (v *veggieDelightBuilder) setSauces() {
	v.sub.sauces = []string{"south west"}
}

func (v *veggieDelightBuilder) getSub() sub {
	return v.sub
}
