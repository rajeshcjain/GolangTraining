package main

import "fmt"

type square struct {
	size float64
}

func (s square) area() float64 {
	return s.size * s.size
}

//We don't mention func in the interfaces
type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(s.area())
}

func main() {
	s := square{64.56}

	/*
	   Here comes the beauty of interface...we can pass any struct which has the implementation of
	   same area as being defined by shape interface..in this case it is square...tomorrow we can
	   have circle,triangle etc.
	*/
	info(s)
}
