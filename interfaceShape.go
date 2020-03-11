package main

import (
	"fmt"
	"math"
)

type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

type Circle struct {
	Centerx float64
	Centery float64
	Radius  float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

func TotalArea(s []Shape) float64 {

	total := 0.0

	for _, currShape := range s {
		total += currShape.Area()
	}

	return total
}

func main() {

	// Remember : always use pointers when internal details change
	s := &Square{5}
	c := &Circle{3, 4, 5}
	shapeSlice := []Shape{s, c}
	total := TotalArea(shapeSlice)
	fmt.Printf("Total area for all shapes in the slice : %f\n", total)

}
