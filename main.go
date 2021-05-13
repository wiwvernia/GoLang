package main

import (
	"fmt"
)

var card = "Ace of Spades"

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func main() {
	cards := newDeck()
	// hand, remainingCards := deal(cards, 2)
	// hand.print()
	// remainingCards.print()
	ca := cards.newDeckFromFile("test.txt")
	// fmt.Println(ca)
	fmt.Println(ca)
	// if err := cards.saveFile("test.txt"); err != nil {
	// 	fmt.Println("Error : ", err)
	// }
	// cards.print()
}
