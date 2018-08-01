package main

import "fmt"

func main() {
	//if we put it after the println statement then it will give and error as undefined.
	x := 12
	fmt.Println(x)
	fmt.Println(y)

}

//Need to see why cant we set it as y := 21...it gives error
//So as we are declaring it in package scope and main function is clouse in the package scope.So it is accessible in main function
//where the order does not matter.

/*
As outer scope that means y(as it has a package level scope) so it is un-clousing the main
function so it is accessible to main function as well.

*/
var y int = 21
