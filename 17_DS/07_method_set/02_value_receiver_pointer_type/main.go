package main

import (
	"math"
	"fmt"
)

type shape interface{
	area() float64
}

type circle struct{
	redius float64
}

func (c circle) area() float64{
	return math.Pi*c.redius*c.redius
}

func info(s shape){
	fmt.Println(s.area())
}

func main(){
	c := circle{23.87}
	info(&c)
}
