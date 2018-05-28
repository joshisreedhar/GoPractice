package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
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

func shuffle(d deck) deck {
	shuffledDeck := d
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for index := range d {
		randomIndex := r.Intn(len(d) - 1)
		shuffledDeck[index], shuffledDeck[randomIndex] = shuffledDeck[randomIndex], shuffledDeck[index]
	}
	return shuffledDeck
}

func save(d deck, f string) {
	ioutil.WriteFile(f, []byte(d.toString()), 0644)
}

func main() {}
