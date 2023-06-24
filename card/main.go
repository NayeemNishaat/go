package main

func main() {
	// var card string = "Ace of Spades" // Note: Initialize new variable

	// card := "Ace of Spades" // Note: Initialize new variable
	// card = "Five of Diamonds" // Note: Reassign value

	// card := newCard()

	// cards := deck{"Ace of Diamonds", newCard()}

	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards := newDeck()
	cards.print()

	// fmt.Println(cards)
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
