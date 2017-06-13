package main

import "fmt"

/*
closure helps us limit the scope of variables by multiple functions
 without closure,for two ofr more function,The scope of the variable
 should be package scope.
*/

func main() {

	x := 0
	increment := func() int {
		x++
		fmt.Println(x)
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
