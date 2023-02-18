package main

import (
	"fmt"
	"math"
)

//Square Struct
type Square struct {
	Length float64 
}

//Area: return the area of square
func (s Square) Area() float64 {
	return s.Length * s.Length
}

//Circle Struct
type Circle struct {
	Radius float64
}

//Circle struct Area method 
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

//sumArea: returns the sum of all the aareas in the slice 
//Shape will be the interface 
func sumArea(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

//Interface: defines a set of methods that a type should implement 
	//In order case, the shape has a single method an area which takes in 
	//no arguments and returns a float64 since both Circle and Sqaure has an Area method
type Shape interface {
	Area() float64
}

func main() {
	s := Square{20}
	c := Circle{10}

	fmt.Println(s.Area())
	fmt.Println(c.Area())

	shapes := []Shape{s,c}

	fmt.Println(sumArea(shapes))

}