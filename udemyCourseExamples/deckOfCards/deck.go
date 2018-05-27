package main

import (
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"♠", "♥", "♣", "♦"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+suit)
		}
	}
	return cards
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func main() {}
