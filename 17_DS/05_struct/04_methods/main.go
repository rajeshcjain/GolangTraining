package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

/*
Below "(p person)" after func is known as receiver and it attaches the function with struct person
*/
func (p person) fullName() string{
	return p.first + p.last
}

func main(){
	p1 := person{"r1","j1",32}
	p2 := person{"r2","j2",21}

	//This is a first way of calling it.
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	//This is a 2nd way of calling it.
	fmt.Println(person.fullName(p1))
}
