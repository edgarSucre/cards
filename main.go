package main

func main() {
	cards := newDeck()
	hand, remainingCards := cards.deal(3)

	remainingCards.print()
	hand.print()
}
