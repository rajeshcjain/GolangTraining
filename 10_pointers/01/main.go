package main

import "fmt"

func main(){

	a := 24

	fmt.Println("The value of a ",a)
	fmt.Println("The address of a ",&a)

	var p *int = &a
	fmt.Println("The value of p ", p)
	fmt.Println("The dererence of p ", *p)

	*p = 42 // p says, that change the value at the memory address, which is stored in p
	fmt.Println("The value of a ",a)

}

//We can pass the memory value instead of a bunch of values(we can pass by reference)
//and we can still change the value of whatever is stored at that memory address
// So this way , we don't pass lot of values  rather we just send the addresses

//Everything in GO is PASS BY VALUE.so when we are passing the address then we are passing the value which is a address.