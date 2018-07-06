package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

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

func (instance deck) print() {
	for index, card := range instance {
		fmt.Println(index, card)
	}
}

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (instance deck) toString() string {
	return strings.Join([]string(instance), ",")
}

func (instance deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(instance.toString()),0666)
}

func newDeckFromFile(filename string) deck  {
	bytes, err :=ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bytes), ","))
}

func (instance deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	for index := range instance {
		newPosition := rand.Intn(len(instance) - 1)
		instance[index], instance[newPosition] = instance[newPosition], instance[index]
	}
}