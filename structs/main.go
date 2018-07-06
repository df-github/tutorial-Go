package main

import "fmt"

// custom type
type contactInfo struct {
	email string
	zipCode int
}

// embedded custom type
type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main()  {
	derrick := person{
		firstName: "Derrick",
		lastName: "Fiedler",
		contact: contactInfo{
			email: "derrick.fiedler@test.de",
			zipCode: 12345,
		},
	}

	fmt.Printf("%+v", derrick)
}
