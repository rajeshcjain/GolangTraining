package main

import "fmt"

type Dog struct{

}

type Cat struct {

}

func main(){

        fmt.Println("Dog....")
	d:= &Dog{}
	d.say()

	fmt.Println("Cat....")
	c:= &Cat{}
	c.say()
}


func (c Cat) say(){
	fmt.Println("I am Cat")
}

func (d Dog) say(){
	fmt.Println("barking")
}



