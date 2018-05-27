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

func Test_DeckShouldPrintItSelf(t *testing.T) {
	stringCards := "A♠, 2♠, 3♠, 4♠, 5♠, 6♠, 7♠, 8♠, 9♠, 10♠, J♠, Q♠, K♠, A♥, 2♥, 3♥, 4♥, 5♥, 6♥, 7♥, 8♥, 9♥, 10♥, J♥, Q♥, K♥, A♣, 2♣, 3♣, 4♣, 5♣, 6♣, 7♣, 8♣, 9♣, 10♣, J♣, Q♣, K♣, A♦, 2♦, 3♦, 4♦, 5♦, 6♦, 7♦, 8♦, 9♦, 10♦, J♦, Q♦, K♦"

	cards := newDeck()

	if cards.toString() != stringCards {
		t.Errorf("Cards are wrongly represented as string. Expected %v got %v", stringCards, cards.toString())
	}
}

func Test_ShouldBeAbleToDealAHandFromDeck(t *testing.T) {

	cards := newDeck()
	handSize := 6
	deal, deck := deal(cards, handSize)

	if len(deal) != handSize {
		t.Errorf("Invalid deal hand size, Expected %v got %v", handSize, len(deal))
	}

	if len(deck) != len(cards)-len(deal) {
		t.Errorf("invalid deck size, deal was not cut")
	}
}
