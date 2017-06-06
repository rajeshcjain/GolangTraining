package main

import "fmt"

//Creating you first type...

//This is a aggregating type...This is encapsulation.It contains fields.
type person struct{
	firstName string
	lastName string
	age int
}

func main(){
	p1 := person{"R1","J1",15000}
	p2 := person{"R2","J2",15000}

	fmt.Println(p1)
	fmt.Println(p2)
}
