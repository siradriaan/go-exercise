package main

func main() {
	// cards := newDeckFromFile("my")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
