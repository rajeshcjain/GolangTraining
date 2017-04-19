package main

import "fmt"

func main(){

	a := 10
	fmt.Println("The value of a is ",a)
	fmt.Println("The address of a is ",&a)

	var b *int = &a

	fmt.Println("The value of b is ",b)
	fmt.Println("Accessing the a with b", *b)
	fmt.Println("Address of b", &b)

}
