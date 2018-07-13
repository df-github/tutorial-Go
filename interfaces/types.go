package main

import "fmt"

type Shouter interface {
	shout()
}

type Dog struct {
	Name string
}
type Lion struct { }
type Cat struct { }


func (Dog) shout(){
	fmt.Println("Wow")
}

func (Lion) shout()  {
	fmt.Println("RoaR")
}

func (Cat) shout()  {
	fmt.Println("Miaow")
}