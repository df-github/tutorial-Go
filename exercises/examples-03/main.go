package main

import "fmt"

func main() {

	derrick := person{
		firstName: "Derrick",
		lastName: "Fiedler",
		contact: contact{
			email: "derrick.fiedler@test.de",
			zipCode: 38530,
		},
	}

	fmt.Printf("%+v", derrick)
}
