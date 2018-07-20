package main

import "fmt"

type Shape interface {
	getArea() float64
}
type Square struct {
	sideLength float64
}
type Triangle struct {
	height float64
	base   float64
}

func (instance *Square) getArea() float64 {
	return instance.sideLength * instance.sideLength
}

func (instance *Triangle) getArea() float64 {
	return 0.5 * instance.height * instance.base
}

func printArea(shape Shape) {
	fmt.Println(shape.getArea())
}

func main() {
	square := &Square{
		sideLength: 4,
	}

	triangle := &Triangle{
		height: 3,
		base:   5,
	}

	printArea(square)
	printArea(triangle)
}
