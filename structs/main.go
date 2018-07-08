package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func (instance person) print() {
	fmt.Printf("%+v\n", instance)
}

func (instance *person) updateName(newFirstName string)  {
	(*instance).firstName = newFirstName
}

func main()  {
	derrick := person{
		firstName: "Derrick",
		lastName: "Fiedler",
		contact: contactInfo{
			email: "derrick.fiedler@test.de",
			zipCode: 38530,
		},
	}

	derrick.updateName("Claude")
	derrick.print()
}
