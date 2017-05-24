package main

import "fmt"

func main(){

	//declaring the array of type in of size 60
	var x [60]int

	fmt.Println(x)
	x[0] = 1
	//This is the function for getting the length of array
	fmt.Println(len(x))
	fmt.Println(x[12])
	fmt.Println(x[0])

}
