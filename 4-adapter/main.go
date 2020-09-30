package main

import "fmt"

func main() {
	espanol := spanishSpeaker{}
	englishwoman := englishSpeaker{}

	fmt.Println("Without translation:")
	fmt.Println("Español says: ", espanol.greetInSpanish())
	fmt.Println("English Woman says: ", englishwoman.greetInEnglish())

	adaptedEnglishwoman := englishToSpanishAdapter{
		speaker: englishwoman,
	}

	fmt.Println("------------")
	fmt.Println("With translation:")
	fmt.Println("Español says: ", espanol.greetInSpanish())
	fmt.Println("English Woman says: ", adaptedEnglishwoman.greetInSpanish())
}
