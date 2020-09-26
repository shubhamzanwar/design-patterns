package main

import (
	"fmt"
	"strings"
)

func describeSub(sub sub) {
	fmt.Printf("bread: %s, cheese: %t, toppings: %s, sauces: %s\n", sub.bread, sub.hasCheese, strings.Join(sub.toppings, ", "), strings.Join(sub.sauces, ", "))
}

func main() {
	veggieDelight := &veggieDelightBuilder{}
	director := &director{
		builder: veggieDelight,
	}
	veggieDelightSub := director.buildSub()
	describeSub(veggieDelightSub)
	fmt.Println("------------")

	director.setBuilder(&chickenTeriyakiBuilder{})
	chickenTeriyakiSub := director.buildSub()

	describeSub(chickenTeriyakiSub)
}
