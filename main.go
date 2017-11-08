package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	cards := newDecFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
