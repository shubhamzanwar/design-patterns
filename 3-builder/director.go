package main

type director struct {
	builder iSubBuilder
}

func (d *director) setBuilder(builder iSubBuilder) {
	d.builder = builder
}

func (d *director) buildSub() sub {
	d.builder.setBread()
	d.builder.setCheese()
	d.builder.setToppings()
	d.builder.setSauces()

	return d.builder.getSub()
}
