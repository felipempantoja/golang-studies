package main

func main() {
	// cards := deck{"aa", "bb"}
	// cards = cards.append("New Card")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingCards := cards.deal(5)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(hand.toString())
	// hand.saveToFile("deck")

	// fmt.Println("--- from file ---")

	// deckFromFile := readFromFile("deck")
	// deckFromFile.print()
}
