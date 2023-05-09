package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("mycards.txt")
	cards := newDeckFromFile("mycards.txt")
	cards.print()
}
