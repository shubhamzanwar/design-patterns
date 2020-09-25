package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func describePizza(pizza iPizza) {
	fmt.Printf("the pizza %s has toppings %s. It costs Rs. %.2f\n", pizza.GetName(), strings.Join(pizza.GetToppings(), ", "), pizza.GetPrice())
}

func describeGarlicBread(garlicBread iGarlicBread) {
	fmt.Printf("the garlic bread, %s costs Rs. %.2f\n", garlicBread.GetName(), garlicBread.GetPrice())
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Dominos or PizzaHut? (D/P)")
	pizzaType, _ := reader.ReadString('\n')
	pizzaType = strings.Split(pizzaType, "\n")[0]

	pizzaFactory, _ := getFactory(pizzaType)

	pizza := pizzaFactory.createPizza()
	garlicBread := pizzaFactory.createGarlicBread()

	describePizza(pizza)
	describeGarlicBread(garlicBread)
}
