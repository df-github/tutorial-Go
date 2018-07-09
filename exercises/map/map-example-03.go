package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "rouge",
		"green": "green",
		"white": "blanc",
	}

	printMap(colors)
}

func printMap(colors map[string]string) {
	for key, value := range colors {
		fmt.Println("The color", key, "translated in french is", value)
	}
}
