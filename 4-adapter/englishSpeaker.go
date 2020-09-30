package main

type englishSpeaker struct{}

func (e *englishSpeaker) greetInEnglish() string {
	return "Hello there"
}
