package main

import (
	"fmt"
)

type Bot interface {
	getGreeting() string
}

type EnglischBot struct {}
type SpanishBot struct {}

func (EnglischBot) getGreeting() string{
	return "Hello There"
}
func (SpanishBot) getGreeting() string{
	return "Halo!"
}

func PrintGreeting(bot Bot){
	fmt.Println(bot.getGreeting())
}

func main() {
	eb := EnglischBot{}
	sb := SpanishBot{}

	PrintGreeting(eb)
	PrintGreeting(sb)

	cat := Cat{}
	lion := Lion{}
	dog := Dog{
		Name:"Wolf",
	}

	fmt.Println("-----------")
	cat.shout()
	lion.shout()
	dog.shout()
}
