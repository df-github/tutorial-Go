package main

import "fmt"

func main()  {
	numbers := []int{0,1,2,3,4,5,6,7,8,9,10}
	for _, number := range numbers {
		if isEven(number){
			fmt.Printf("%v is even \n", number)
		} else {
			fmt.Printf("%v is odd \n", number)
		}
	}
}

func isEven(number int) bool {
	return number % 2 == 0
}
