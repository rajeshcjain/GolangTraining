package main

import "fmt"

func main(){

	customNumber := make([]int,3)
	customNumber[0] = 12
	customNumber[1] = 21
	customNumber[2] = 31

	//Here we can only access till index 2 as there are there is not array to access
	//beyond this point.We need to call append function to increase the length of the
	// underlying array.
	fmt.Println(customNumber[0])
	fmt.Println(customNumber[1])
	fmt.Println(customNumber[2])

	greetings := make([]string,3,3)
	greetings[0] = "R1"
	greetings[1] = "S1"
	greetings[2] = "A1"

	fmt.Println(greetings[0])
	fmt.Println(greetings[1])
	fmt.Println(greetings[2])


}
