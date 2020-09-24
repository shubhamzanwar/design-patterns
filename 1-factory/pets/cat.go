package pets

type cat struct {
	name  string
	sound string
	age   int
}

func (c cat) GetName() string {
	return c.name
}

func (c cat) GetAge() int {
	return c.age
}

func (c cat) GetSound() string {
	return c.sound
}
