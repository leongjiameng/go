package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	area := ((22.0 / 7.0) * math.Pow(c.radius, 2))
	return area
}

func (s square) area() float64 {
	area := math.Pow(s.length, 2)
	return area
}

func info(s shape) {
	switch s.(type) {
	case square:
		fmt.Println("My length is", s.(square).length, "and my area is", s.(square).area())
	case circle:
		fmt.Println("My radius is", s.(circle).radius, "and my area is", s.(circle).area())
	}
}

func main() {
	aCircle := circle{
		radius: 15,
	}

	aSquare := square{
		length: 15,
	}

	info(aCircle)
	info(aSquare)

}

/*
Hands-on exercise #5
● create a type SQUARE
● create a type CIRCLE
● attach a method to each that calculates AREA and returns it
○ circle area= π r 2
○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
*/
