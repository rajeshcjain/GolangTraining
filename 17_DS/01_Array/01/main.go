package main


import "fmt"

func main(){

	//This is how we are creating the Array
	var x[60] string
	//To create a slice you can create it using var x[] string...so if we remove the length from
	//The braces then its a slice
	/*
	               IMPORTANT
	               Arrays are not dynamic but slices are dynamic

	So when i am declaring it with string then it will be initializing with blank....if the type is int
	then it will be initiazed with ZERO

	*/
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 60;i < 120;i++{
		x[i-60] = string(i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

}
