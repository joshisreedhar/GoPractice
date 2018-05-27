package main

import (
	"strings"
	"testing"
)

func Test_DeckSize(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected 52 cards, found %v", len(cards))
	}
}

func Test_DeckShouldNotHaveDuplicateCard(t *testing.T) {
	checkedCards := []string{}
	cards := newDeck()
	for _, card := range cards {
		if contains(checkedCards, card) {
			t.Errorf("A duplicate card found")
		}
		checkedCards = append(checkedCards, card)
	}
}

func Test_ASuitShouldHave13Cards(t *testing.T) {

	suits := []string{"♠", "♥", "♣", "♦"}
	cards := newDeck()
	suitCount := 0
	for _, suit := range suits {

		for _, card := range cards {
			if strings.HasSuffix(card, suit) {
				suitCount = suitCount + 1
			}
		}

		if suitCount != 13 {
			t.Errorf("%v does not have 13 cards", suit)
		}
		suitCount = 0
	}
}
