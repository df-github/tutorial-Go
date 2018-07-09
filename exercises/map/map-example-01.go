package main

import "fmt"

/**
  first way how to declare and initialize a map in Go.
 */
func main()  {
	colors := map[string]string{
		"red": "rouge",
		"green": "vert",
		"white": "blanc",
	}

	fmt.Println(colors)
}
