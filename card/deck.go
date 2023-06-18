package main

// Point: Create a new type of deck (slice of string)
type deck []string

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}
