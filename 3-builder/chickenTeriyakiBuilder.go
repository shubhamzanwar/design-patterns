package main

type chickenTeriyakiBuilder struct {
	sub
}

func (c *chickenTeriyakiBuilder) setBread() {
	c.sub.bread = "italian"
}

func (c *chickenTeriyakiBuilder) setCheese() {
	c.sub.hasCheese = true
}

func (c *chickenTeriyakiBuilder) setToppings() {
	c.sub.toppings = []string{"roasted chicken", "olives", "onions", "jalapeños"}
}

func (c *chickenTeriyakiBuilder) setSauces() {
	c.sub.sauces = []string{"chilli", "bbq"}
}

func (c *chickenTeriyakiBuilder) getSub() sub {
	return c.sub
}
