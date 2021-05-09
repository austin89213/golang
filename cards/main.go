package main

import "fmt"

func mian() {
	card := newCard()

	fmt.Println(card)
}

func newCard() {
	return "Five of Diamonds"
}
