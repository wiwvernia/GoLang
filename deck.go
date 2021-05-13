package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamond", "Heart", "Clubs"}
	cardNumber := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nice", "Ten", "Jack", "Queen", "King"}
	for _, suits := range cardSuits {
		for _, cNumber := range cardNumber {
			nametest := cNumber + " of " + suits
			cards = append(cards, nametest)
		}
	}
	return cards
}

func (d deck) print() {
	fmt.Println(d)
}

func (d deck) getTitle() deck {
	return d
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0777)
}

func (d deck) newDeckFromFile(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	return strings.Split(string(content), ", ")
}
