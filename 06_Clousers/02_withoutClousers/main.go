package main

import "fmt"

var x = 0

// Now for x,it is accessible to all the functions as it has a package level scope.
func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}
