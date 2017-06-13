package main

import "fmt"

func main() {

	mytSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T \n", mytSlice)

	for i, val := range mytSlice {
		fmt.Println("index : ", i)
		fmt.Println("value : ", val)
	}
}
