package main

func main() {
	cards := newDeck()
	cards = append(cards, "Six of Spades")

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
