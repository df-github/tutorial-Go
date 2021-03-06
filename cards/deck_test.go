package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 12 {
		t.Errorf("Expected deck length of 16 but got %v", len(deck))
	}

	firstElement := deck[0]
	expectedFirstElement := "Ace of Spades"
	if firstElement != expectedFirstElement {
		t.Errorf("Expected first element to be '%v' but got %v", expectedFirstElement, firstElement)
	}

	lastElement := deck[len(deck) -1]
	expectedLastElement := "Three of Clubs"
	if lastElement != expectedLastElement {
		t.Errorf("Expected last element to be '%v' but got %v", expectedLastElement, lastElement)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T){
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeckFromFile := newDeckFromFile(filename)

	if len(loadedDeckFromFile) != len(deck) {
		t.Errorf("Expected deck length of %v, but got %v", len(deck), len(loadedDeckFromFile))
	}

	os.Remove(filename)
}

func TestDeal(t *testing.T) {
	deck := newDeck()

	handSize := 5
	handDeck, restDeck := deal(deck, handSize)

	if handSize != len(handDeck) {
		t.Errorf("Expected of having %v cards, but got %v", handSize, len(handDeck))
	}

	restCardsSize := len(deck) - handSize
	if restCardsSize != len(restDeck){
		t.Errorf("Expected of having rest %v cards, but got %v", restCardsSize, len(restDeck))
	}
}