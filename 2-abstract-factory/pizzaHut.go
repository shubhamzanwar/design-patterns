package main

type pizzaHutFactory struct{}

type pizzaHutPizza struct {
	pizza
}

type pizzaHutGarlicBread struct {
	garlicBread
}

func (p *pizzaHutFactory) createPizza() iPizza {
	return &pizzaHutPizza{
		pizza{
			name:     "pepperoni",
			price:    230.3,
			toppings: []string{"olives", "mozzarella", "pork"},
		},
	}
}

func (p *pizzaHutFactory) createGarlicBread() iGarlicBread {
	return &pizzaHutGarlicBread{
		garlicBread{
			name:  "garlic bread",
			price: 180.99,
		},
	}
}
