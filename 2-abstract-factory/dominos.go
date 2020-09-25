package main

type dominosFactory struct{}

type dominosPizza struct {
	pizza
}

type dominosGarlicBread struct {
	garlicBread
}

func (d *dominosFactory) createPizza() iPizza {
	return &dominosPizza{
		pizza{
			name:     "margherita",
			price:    200.5,
			toppings: []string{"tomatoes", "basil", "olive oil"},
		},
	}
}

func (d *dominosFactory) createGarlicBread() iGarlicBread {
	return &dominosGarlicBread{
		garlicBread{
			name:  "cheesy bread sticks",
			price: 150.00,
		},
	}
}
