package main

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

func main() {}
