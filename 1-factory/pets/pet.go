package pets

type pet struct {
	name  string
	age   int
	sound string
}

func (p *pet) GetName() string {
	return p.name
}

func (p *pet) GetSound() string {
	return p.sound
}

func (p *pet) GetAge() int {
	return p.age
}
