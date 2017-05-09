package main

import "fmt"


// this is interesting here.Here we have a function which
//returns the anonymous function which returns the string
//and we are returning the anonymous function which returns string "rajesh jain"
func makegreeter() func() string{
	return func() string{
		return "Rajesh Jain"
	}
}

func main(){
	greeter := makegreeter()
	fmt.Println(greeter())
	fmt.Printf("%T\n",greeter)
}
