package main

import "fmt"

var x = 0

// Now for x,it is accessible to all the functions as it has a package level scope.
// and the below function is a part of the package level scope so it can access it.
func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}
