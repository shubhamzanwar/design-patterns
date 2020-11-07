package main

import "fmt"

type voiceAssistant struct {
	next step
}

func (v *voiceAssistant) run(cust *customer) {
	fmt.Println("[Voice Assistant] Serving the customer: ", cust.name)
	v.next.run(cust)
}

func (v *voiceAssistant) setNextStep(next step) {
	v.next = next
}

type associate struct {
	next step
}

func (a *associate) run(cust *customer) {
	if cust.isHighPriority {
		fmt.Println("Redirecting customer directly to manager")
		a.next.run(cust)
		return
	}
	fmt.Println("[Associate] Serving the customer: ", cust.name)
	a.next.run(cust)
}

func (a *associate) setNextStep(next step) {
	a.next = next
}

type manager struct {
	next step
}

func (a *manager) run(cust *customer) {
	fmt.Println("[Manager] Serving the customer: ", cust.name)
}

func (a *manager) setNextStep(next step) {
	a.next = next
}
