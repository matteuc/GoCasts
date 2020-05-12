package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func (sq square) getArea() float64 {
	return math.Pow(sq.sideLength, 2.0)
}

func printArea(s shape) {
	fmt.Printf("Area: %v\n", s.getArea())
}

func main() {
	t := triangle{height: 16, base: 4}
	sq := square{sideLength: 4}

	printArea(t)
	printArea(sq)

}
