package main

import "fmt"

type Person struct{
	first string
	last string
	age int
}

//So here we are overriding the person's first field
type DoubleZero struct{
	Person
	first string
	LicenceToKill bool
}

func main(){

	p1 := DoubleZero{
		Person : Person{
			first : "James",
			last : "Bond",
			age : 35,
		},
		first: "I have License to kill",
		LicenceToKill:true,
	}

	//But we can access the Inner value as well outer value
	fmt.Println("full name ",p1.Person.first + p1.Person.last, " Outer first name ", p1.first)


}
