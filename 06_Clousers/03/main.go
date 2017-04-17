package main

import "fmt"

func main(){
	var x = 0

	//These are called anonymous functions.These are functions without name.These are func expressions assigning a func to a variable
        increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

}
