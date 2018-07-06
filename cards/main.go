package main

func main()  {
	filename := "my-cards"

	cards := newDeck()
	cards.saveToFile(filename)

	cardsFromFile := newDeckFromFile(filename)
	cardsFromFile.shuffle()
	cardsFromFile.print()
}


