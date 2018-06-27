package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue + " of " + cardSuite)
		}
	}

	return cards
}

func (instance deck) print()  {
	for index, card := range instance {
		fmt.Println(index, card)
	}
}