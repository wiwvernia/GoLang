package main

var card = "Ace of Spades"

func main() {
	cards := deck{}
	// var card string = "Ace of Spades"
	// for _, card := range cards {
	// 	fmt.Println(card)
	// 	cards.print()
	// }

	cardSuits := []string{"Spades", "Diamond", "Heart", "Clubs"}
	cardNumber := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nice", "Ten", "Jack", "Queen", "King"}
	for _, suits := range cardSuits {
		for _, cNumber := range cardNumber {
			nametest := cNumber + " of " + suits
			cards = append(cards, nametest)
		}
	}
	cards.print()
	// fmt.Println(cards)

}

func newCard() string {
	return "2 of Spades"
}
