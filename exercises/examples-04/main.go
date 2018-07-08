package main

import "fmt"

/**
 slice is a reference Type. That is why we dont care about pointer.
 */
func main()  {
	mySlices := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlices)

	fmt.Println(mySlices)
}

func updateSlice(slices []string) {
	slices[0] = "Bye"
}

