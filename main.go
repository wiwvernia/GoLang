package main

var card = "Ace of Spades"

func main() {
	cards := deck{}
	// var card string = "Ace of Spades"
	cards = append(cards, "3 of Spades")
	// for _, card := range cards {
	// 	fmt.Println(card)
	// 	cards.print()
	// }

	cardSuits := []string{"Spades", "Diamond", "Heart"}
	cardNumber := []string{"1", "2", "3", "4"}
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
