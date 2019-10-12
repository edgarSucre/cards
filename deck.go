package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

//	This sintaxis means that any variable of type deck has access to the print method
func (d deck) print() {

	//range makes a for a foreach
	for _, card := range d {
		fmt.Println(card)
	}
}

// deal (n) cards from the deck, and return the remaining cards in the deck
func deal(d deck, n int) (deck, deck) {
	return d[:n], d[n:]
}
