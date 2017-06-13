package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	/*
		Here we are learning a important concept with if condition.Here first we get
		the value and whether that value exists in the map or not.Based on that
		exists value will be set to true /false.Then after the ";" we have a condition
		which tells whether its true or false.


	*/
	if val, exists := myGreeting[1]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)

	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)

}
