package main

import "fmt"

/*
I had a confusion about..if car can be a cloth and asked todd about it...below is the response

Values can be of several types in Go.

So you have a value of type "Car" assigned to the variable c

Any value of type "Car" also has access to a method color()

You also have two interfaces defined

Read the definition of those interface types

You have created two new TYPES - "Vehicle" and "Cloth"

Remember, a VALUE is of a certain TYPE

With x := 7 we have the VALUE 7 and it is of type INT

So back to the interfaces

An interface says ,"Hey baby, if you have these methods, then you're my type"

That's a funny way of saying what interfaces do - and it's also true.

So the interface says which METHODS a VALUE must have to IMPLICITLY IMPLEMENT the interface

With the TYPES "Vehicle" and "Cloth", they are each saying that a VALUE with a method "Color()"

will implement each of those interfaces.

So you have a value of type "Car" assigned to the variable c

Any value of type "Car" also has access to a method color()

So the variable c is THREE TYPES - "Car" "Vehicle" and "Cloth"

Therefore, anywhere a VALUE of TYPE "Car" "Vehicle" or "Cloth" is asked for, the variable "c" can be used (because it is all three of those types)

I hope this helps!



*/

/*
Interface is a type that just declare the behavior...it does not implement it.Its the user defined types
that does it...so we define the struct here and it implements the method.
*/
type Vehicle interface {
	color() string
}

type Cloth interface {
	color() string
}

type Car struct {
	wheels       int
	AirCondition bool
}

func (c Car) color() string {
	return "black"
}

func info(c Cloth) {
	fmt.Println(c.color())
}

func main() {
	c := Car{4, true}
	info(c)
}
