package main

import "fmt"

func main(){
	derrick := person{
		firstName: "Derrick",
		lastName: "Fiedler",
	}

	log(derrick)

	var paul person
	paul.firstName = "Paul"
	paul.lastName = "Meyer"

	log(paul)
}

func log(object interface{})  {
	fmt.Printf("%+v \n", object)
}
