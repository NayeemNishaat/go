package main

import "fmt"

type bot interface {
	getGreeting() string
	// doStuff(string, int) (string, error) // Important: Need to define all the methods to qualify as bot
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	eb.ebSpecific()
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Note: Custome logic for English greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// Note: Custome logic for Spanish greeting
	return "Hola!"
}

func (englishBot) ebSpecific() {
	fmt.Println("EB Specific")
}
