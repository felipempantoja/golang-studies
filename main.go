package main

import "fmt"

var deckSize int

func main() {
	// var card string = "Ace os Spades"
	// var card = "Ace os Spades"
	// card := newCard()
	// fmt.Println(card)
	// deckSize = 40

	// printState()

	cards := []string{newCard(), newCard()}
	fmt.Println(cards)

	cards = append(cards, "New Card")
	// fmt.Println(newCards)

	for i, card := range cards {
		fmt.Println(i, card)
	}

	for i := 0; i < len(cards); i++ {
		fmt.Println(i, cards[i])
	}
}

func newCard() string {
	return "Five of Diamonds"
}
