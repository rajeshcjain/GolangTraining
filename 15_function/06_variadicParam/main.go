package main

import (
	"fmt"
)

func main() {

	//Here we are passing the slice which is nothing but a list.
	//when we pass slice to function we don't need to type cast it
	//as these are constants and when we pass it to function
	//which accepts the variadic of int it will convert it in to
	//type int.
	//In Go, when have constants then those constants does not have any type..
	//we call that these constants are of a kind i.e they are of there own kind
	//which are not assigned any type..so they can be assigned any type..in this case
	//these does not have any type and hence can be assigned a type(in this case int) when passed to a function
	//This way we need not to type cast it.
	fmt.Println(greet(1, 2, 3, 4, 5, 6, 7))
}

//These functions are called variadic functions which accepts variadic parameter.
func greet(sf ...int) int {

	//Now when we print the sf it will print the slice which is a list.
	fmt.Println(sf)

	//Here it will print the type of the slice which is []int
	fmt.Printf("%T \n", sf)

	total := 0

	// here range operator returns the index and value at that index.
	//so we are ignoring the index and keeping the value and then
	// adding it total.
	for _, v := range sf {
		total += v
	}
	return total
}
