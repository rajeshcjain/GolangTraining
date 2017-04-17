package main

import "fmt"

func main() {

	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "I am in inner scope"
		fmt.Println(y)
	}
	//The below statement will give error as undefined as it is out of scope.
	// fmt.Println(y)

}
