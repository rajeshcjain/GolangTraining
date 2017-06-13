package main

import (
	"fmt"

	"math"
)

//const  pi  = 3.14

type square struct {
	size float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

//So here we are implementing the shape interface
func (s square) area() float64 {
	return s.size * s.size
}

//Here we are implementing the circle interface
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func info(s shape) {
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(s.area())
}

func main() {
	s := square{12}
	info(s)
	c := circle{12}
	info(c)
}
