package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	newCards := newDeckFromFile("my_cards")
	newCards.shuffle()
	newCards.print()
}
