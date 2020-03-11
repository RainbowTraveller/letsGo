package main

import (
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length int
}

func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func NewSquare(x int, y int, side int) (*Square, error) {
	if side < 0 {
		return nil, fmt.Errorf("Side can not be -ve")
	}
	s := &Square{
		Center: Point{x, y},
		Length: side,
	}
	return s, nil
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	pt := Point{10, 20}
	pt.Move(2, 5)
	fmt.Println(pt)
	sq, err := NewSquare(2, 5, 4)
	if err != nil {
		log.Fatalf("ERROR : can not instantiate a square")
	}

	fmt.Printf("Original Square : %v :\n", sq)
	sq.Move(3, 3)
	fmt.Printf("Square Moved to : %v :\n", sq)
	fmt.Println(sq.Area())
}
