package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

type DoubleZero struct {
	person
	LicenseToKill bool
}

func (p person) greetings(){
	fmt.Println(p.name, " : ",p.age)
}

func (d DoubleZero) greetings(){
	fmt.Println(d.person.name," : ",d.person.age, " : ", d.LicenseToKill)
}

func main(){

	p2 := person{
		name : "R1",
		age : 32,
	}

	p1 := DoubleZero{
		person : person{
			name : "James Bond",
			age : 35,
		},
		LicenseToKill: true,
	}

	//This will call DoubleZero greeting method
	p1.greetings()
	//This will call person method
	p2.greetings()

	//This will call the greeting method of the person.But it will call with the
	//name and age of p1...but will call the method of person...so the output will
	// be "James bond : 35"
	p1.person.greetings()
}
