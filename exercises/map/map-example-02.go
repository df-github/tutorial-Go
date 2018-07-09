package main

import "fmt"

func main(){
	//var colors map[string]string

	colors := make(map[string]string)
	log(colors)

	colors["white"] = "blanc"

	fmt.Println(colors)

	delete(colors, "white")
}

func log(colors map[string]string) (int, error) {
	return fmt.Println(colors)
}
