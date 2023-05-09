package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	fmt.Println("-- START OF DECK --")
	for i, card := range d {
		fmt.Println(i+1, card)
	}
	fmt.Println("-- END OF DECK --")
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardNums := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	cardSpecials := []string{"King", "Queen", "Joker"}

	for _, cardSuit := range cardSuits {
		for _, cardNum := range cardNums {
			cards = append(cards, cardNum+" of "+cardSuit)
		}
		for _, cardSpecial := range cardSpecials {
			cards = append(cards, cardSpecial+" of "+cardSuit)
		}
	}

	return cards
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) toString() string {
	deckStringArr := []string(d)
	deckStr := strings.Join(deckStringArr, ", ")
	return deckStr
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	fileBS, bErr := ioutil.ReadFile(filename)

	if bErr != nil {
		fmt.Println("Error:: ", bErr)
		os.Exit(1)
	}

	stringFileBS := strings.Split(string(fileBS), ", ")
	return deck(stringFileBS)
}

func getRandomCard(d deck) string {
	return "Random card"
}
