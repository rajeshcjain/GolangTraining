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

func (c *circle) area() float64{
	return math.Pi*c.redius*c.redius
}

func info(s shape){
	fmt.Println(s.area())
}
/*
Below code will not work as explained below

There are typed and untyped constants

const hello = "hello" \\Its a untyped constant
const typedConst string = "hello" \\Its a typed constant

untyped constant - a constant that does not have any fixed type
 Its a kind.Not yet forced to obey the strict rules that combining
 different typed values.
 It is this notion that makes it possible to use constants in Go with with
 great freedom.

 This is useful,for instance

 42

 So if we don't have untyped constants than we have to do the conversion for
 every type;so we have to convert it in to int using int(42).

 SO because of this...we will not be aware of the c..as the type of it is not known
 yet and if you don't know the type of something then how do you know the address
 of it.

*/
func main(){
	c := circle{23.87}
	info(c)
}
