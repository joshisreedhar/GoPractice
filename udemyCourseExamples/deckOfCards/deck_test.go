package main

import (
	"io/ioutil"
	"os"
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

func Test_ShouldBeAbleToRandomlyShuffleADeck(t *testing.T) {
	cards := newDeck()
	shuffledCards := shuffle(cards)

	if !areCardsShuffled(shuffledCards) {
		t.Errorf("Cards were not shuffled properly, found majority cards with same suit together")
	}

}

func Test_ShouldBeAbleToShuffleADeal(t *testing.T) {
	d := newDeck()

	hand, deck := deal(d, 5)

	shuffledHand := shuffle(hand)

	shuffledDeck := shuffle(deck)

	if !areCardsShuffled(shuffledHand) {
		t.Errorf("Hand not shuffled properly")
	}

	if !areCardsShuffled(shuffledDeck) {
		t.Errorf("Remaining deck not shuffled properly")
	}

}

func Test_CanSaveADeckToFile(t *testing.T) {
	d := newDeck()
	f := "_deckfile"
	save(d, f)
	_, err := ioutil.ReadFile(f)
	if err != nil {
		t.Errorf("could not write to file")
	}
	os.Remove(f)
}

func Test_SavedDeckFileHasRightContent(t *testing.T) {
	d := newDeck()
	f := "_deckfile"
	save(d, f)
	content, err := ioutil.ReadFile(f)
	if err == nil && string(content) != d.toString() {
		t.Errorf("Written file content is not correct")
	}
	os.Remove(f)
}

func areCardsShuffled(d deck) bool {
	cardCount := 0
	immediateCardCount := 0
	cardSuit := ""
	previousCardSuit := ""
	randomLimit := len(d) / 2

	for _, card := range d {

		cardSuit = card[1:]
		cardCount = cardCount + 1
		if cardSuit == previousCardSuit {
			immediateCardCount = immediateCardCount + 1
		} else {
			immediateCardCount = 1
		}

		if cardCount == randomLimit {
			if immediateCardCount == randomLimit {
				return false
			}
			cardCount = 0
		}
	}
	return true
}
