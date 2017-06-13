package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	/*
		Here initial length of the slice is 3 so if you try to acess the 4th index;It will throw and
		error.So to avoid it we have to use append function as discussed in previous tutorials.
	*/

	//greeting[3] = "suprabadham"

	greeting = append(greeting, "suprabadham")

	fmt.Println(greeting[2])
}
