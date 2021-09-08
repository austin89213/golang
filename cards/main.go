package main

import "fmt"

func main() {
	// card := newCard()
	// fmt.Println(card)
	cards := []string{"Ace of Diamonds", newCard()}
	// fmt.Println(cards)
	cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	// Iterate a slice or list

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
