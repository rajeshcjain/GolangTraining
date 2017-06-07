package main

import "fmt"

type Person struct{
	first string
        last string
        age int
}

type DoubleZero struct{
	Person
	LicenceToKill bool
}

func main(){
	//This is a Second way of creating the value of struct.
	p1 := DoubleZero{
		Person: Person{
			first:"R1",
			last:"J1",
			age:32,
		},
		LicenceToKill:true,
	}

	p2 := DoubleZero{
		Person: Person{
			first:"R2",
			last:"J2",
			age:21,
		},
		LicenceToKill:true,
	}

	fmt.Println(p1.Person.first + p1.Person.last)
	fmt.Println(p2.Person.first + p2.Person.last)
	fmt.Println(p1.LicenceToKill)
	fmt.Println(p2.LicenceToKill)


}


