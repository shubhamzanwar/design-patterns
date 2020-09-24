package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./pets"
)

func describePet(pet pets.IPet) string {
	return fmt.Sprintf("%s is %d years old. It's sound is %s", pet.GetName(), pet.GetAge(), pet.GetSound())
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Are you a dog person or cat person? (dog/cat)")
	petType, _ := reader.ReadString('\n')
	petType = strings.Split(petType, "\n")[0]

	pet := pets.PetFactory(petType)
	petDescription := describePet(pet)

	fmt.Println(petDescription)
}
