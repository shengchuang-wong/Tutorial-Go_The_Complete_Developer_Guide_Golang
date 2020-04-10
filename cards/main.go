package main

func main() {
	// var card string = "Ace of Spades" // equivalent to below
	// card := "Ace of Spades"
	// card = "Five of Diamonds" // reassign
	// card := newCard()
	// cards := deck{"Foud of Diamonds", newCard()}
	// cards := newDeck()
	// cards.saveTolFile("king")
	// fmt.Println(cards.toString())
	// cards := newDeckFromFile("my_cards")

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
