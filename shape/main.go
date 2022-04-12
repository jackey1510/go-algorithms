package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	sq := square{sideLength: 5}
	tr := triangle{base: 2, height: 3.5}

	printArea(sq)
	printArea(tr)
}
