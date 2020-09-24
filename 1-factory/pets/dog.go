package pets

type dog struct {
	name  string
	sound string
	age   int
}

func (d dog) GetName() string {
	return d.name
}

func (d dog) GetAge() int {
	return d.age
}

func (d dog) GetSound() string {
	return d.sound
}
