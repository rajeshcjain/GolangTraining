package main

import (
	"fmt"
)

func main() {

	//Here we are declaring the function expression and the calling it.
	//It is also called anonimous function.
	greeting := func() {
		fmt.Println("Hellow world")
	}

	fmt.Printf("%T\n", greeting)
	greeting()
}
