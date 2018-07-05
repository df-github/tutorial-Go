package main

import "fmt"

func main()  {
	cards := newDeck()

	cards.print()

	fmt.Println("-------------------")
	dealedCards, restDeck :=cards.deal(5)

	fmt.Println("-##### Dealed Cards #####-")
	dealedCards.print()

	fmt.Println("-##### Rest Cards #####-")
	restDeck.print()

}


