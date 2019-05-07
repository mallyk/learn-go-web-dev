package main

import (
	"fmt"
	"math"
)

type square struct {
	sideLength float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func info(s shape) {
	fmt.Println("the area is:", s.area())
}

func main() {
	s1 := square{4}
	c1 := circle{3}

	info(s1)
	info(c1)
}
