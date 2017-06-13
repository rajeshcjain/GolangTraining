package main

import "fmt"

//Callback is passing the function as an argument.

//Here we define the slice of type []int as first parameter
//and second param as callback function.Here "callback" is the name
// and here "func(int)" tells is the type so this callback function
// accepts the function which accepts int as a parameter.
func visit(numbers []int, callback func(int)) {
	//range first parameter it returns is index and second is the actual number at
	// that index.so here we are ignoring the first part of the return statement and
	// storing the 2nd part in n.
	for _, n := range numbers {
		callback(n)
	}
}

func main() {

	//we are passing the same type in visit
	visit([]int{1, 2, 3, 4, 5, 6}, func(n int) {
		fmt.Println(n)
	})
}
