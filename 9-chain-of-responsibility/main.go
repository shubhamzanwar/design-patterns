package main

import "fmt"

func main() {
	m := &manager{}

	assoc := &associate{}
	assoc.setNextStep(m)

	va := &voiceAssistant{}
	va.setNextStep(assoc)

	// Chain formation complete

	// Start chain execution for normal customer
	normalCust := &customer{
		name: "Bob",
	}

	va.run(normalCust)

	fmt.Println("===================")

	// Start chain execution for high priority customer
	highPriorityCust := &customer{
		name:           "John",
		isHighPriority: true,
	}

	va.run(highPriorityCust)
}
