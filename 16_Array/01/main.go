package main

import "fmt"

func main() {

	//declaring the array of type int of size 60
	//Here we are creating a array of size...please remember if you are mentioning a number
	//in the bracket then its array and if you don't then its a slice.

	/*

	so var x [60]int is a array as we have mentioned the 60 in the bracket.
	Please note that the type of the array is [60]int and number mentioned in the
	bracket is the part of the type of the array.

	and var x []int = {} is a slice as we have not mentioned the number.

	so remember when you try to pass the array of type [60]int to function greet()
	it will be fine but when you try to change the function length from from 60 -> something else
	then type of the input argument to the function will change; and hence it will throw the
	exception at compile time.

	Please remember that in GO,Arrays are not dynamic..slices are dynamic.. that is why
	slices are used more.

	The size of the array can not be changed as that is the part of the the type of the array.

	So when you declare an array it will be initialized with ZERO and when we call len function it will
	return 60 which is part of the type of the array.

	*/
	var x [60]int

	fmt.Println(x)
	fmt.Printf("%T \n",x)
	x[0] = 1
	//This is the function for getting the length of array
	fmt.Println(len(x))
	fmt.Println(x[12])
	fmt.Println(x[0])

	greet(x)

}


func greet(arr [60]int){
	fmt.Println("This is just for testing")
}
