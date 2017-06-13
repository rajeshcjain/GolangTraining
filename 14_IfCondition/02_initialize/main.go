package main

import "fmt"

func main() {

	b := true

	// here we are initializing the food in the if condition only.
	//There is no need of seprate statement before the
	// if condition.The scope of food is with in the if condition only.
	if food := "Rajesh jain"; b {
		fmt.Println(food)
	}

}
