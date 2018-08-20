package main

import (
	"fmt"
)

func main() {


	//Here we are declaring the function expression and the calling it.
	//It is also called anonymous function.

	//There is one good thing which need to be understand here for function types
	//In go, functions are of one type...and each function belongs to a func() type.
	/*
	ex :-

	the below function and all the function which does not accept any parameters and
	return nothing belongs to one type.

	so function func main(){}, func greet() and func user() belongs to same type and that func()

	*/
	greeting := func() {
		fmt.Println("Hello world")
	}

	/*

	Similarly all the functions which accepts variadic parameters and does not return nothing belongs
	to a type which func(...int)

	*/

	greeting1 := func(sf ...int){
		fmt.Println("The length is ", len(sf))
	}

	/*
		Similarly all the functions which accepts variadic parameters and return int belongs
		to a type which func(...int) int
	*/
	greeting2 := func(sf ...int) int{
		fmt.Println("The length is ", len(sf))
		return len(sf)
	}

	/*
		Similarly all the functions which accepts string and return int belongs
		to a type which func(string) int
	*/

	greeting3 := func(v1 string) int{
		fmt.Println("The length is ", len(v1))
		return len(v1)
	}

	fmt.Printf("%T\n", greeting)
	fmt.Printf("%T\n", greeting1)
	fmt.Printf("%T\n",greeting2)
	fmt.Printf("%T\n",greeting3)
	greeting()
	greeting1()
	greeting2()
}
