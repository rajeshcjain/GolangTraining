package main

import "fmt"

// this is interesting here.Here we have a function which
//returns the anonymous function which returns the string
// and we are returning the anonymous function which returns string "rajesh jain"

//The type of this function is func() func() string as this function which return a function
//which internally return a string.
func makegreeter() func() string {
	return func() string {
		return "Rajesh Jain"
	}
}

func main() {
	greeter := makegreeter()
	fmt.Printf("%T\n",makegreeter)
	fmt.Println(greeter())
	fmt.Printf("%T\n", greeter)
}
