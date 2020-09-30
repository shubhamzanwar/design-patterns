package main

func translate(englishMesage string) string {
	// some complex logic to translate
	return "¡Hola!"
}

type englishToSpanishAdapter struct {
	speaker englishSpeaker
}

func (a *englishToSpanishAdapter) greetInSpanish() string {
	spanishMessage := translate(a.speaker.greetInEnglish())
	return spanishMessage
}
