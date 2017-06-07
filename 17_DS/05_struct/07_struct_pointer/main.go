package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){

	p1 := &person{"R1",32}
	fmt.Println(p1)
	//Go is actually converting it in to "*p1.name"
	fmt.Println(p1.name)
	//Go is actually converting it in to "*p1.age"
	fmt.Println(p1.age)
	fmt.Printf("%T",p1)


}
