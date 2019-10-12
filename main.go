package main

func main() {
	cards := newDeck()
	hand, cards := deal(cards, 3)

	cards.print()
	hand.print()
}
