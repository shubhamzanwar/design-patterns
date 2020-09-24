package pets

// Pet defines the general structure of all pets
type Pet interface {
	GetName() string
	GetSound() string
	GetAge() int
}

// PetFactory is a factory that return the pet requested
func PetFactory(petType string) Pet {
	if petType == "dog" {
		return dog{
			name:  "Chester",
			age:   2,
			sound: "bark",
		}
	} else if petType == "cat" {
		return cat{
			name:  "Mr. Buttons",
			age:   3,
			sound: "meow",
		}
	}
	return nil
}
