package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // Note: Initialize new variable

	// card := "Ace of Spades" // Note: Initialize new variable
	// card = "Five of Diamonds" // Note: Reassign value

	// card := newCard()

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
