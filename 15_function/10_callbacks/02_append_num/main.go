package main

import "fmt"

//We could have done it without using the callbacks but here the scope of it is till this function.
//Other then that we are using anonymous function here.SO no name.
func filter(numbers []int, callbacks func(int) bool) []int {
	//This is how we declare a empty slice.This is important.
	// var appendedItems []int ==> Another way of declaring the below statement
	appendedItems := []int{}

	for _, n := range numbers {
		if callbacks(n) {
			appendedItems = append(appendedItems, n)
		}
	}

	//Below statement should give me "func(int) bool" which proves that
	// function is a type in GoLang
	fmt.Printf("%T\n", callbacks)
	return appendedItems
}

func main() {

	returnedVal := filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n > 1
	})

	fmt.Println(returnedVal)

}
