package main

type contact struct {
	email string
	zipCode int
}

// kompakt way to embed custom typ
type person struct {
	firstName string
	lastName string
	contact
}

