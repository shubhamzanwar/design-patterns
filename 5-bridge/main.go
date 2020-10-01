package main

import "fmt"

func main() {
	maleDeveloper := Male{
		name:       "John",
		age:        22,
		department: Developer{},
	}

	fmt.Println(maleDeveloper.describePerson())
	fmt.Println("-------------")

	femalePM := Female{
		name:       "Natalie",
		age:        24,
		department: PM{},
	}

	fmt.Println(femalePM.describePerson())
}
