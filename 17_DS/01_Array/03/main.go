package main

import "fmt"

func main() {

	var x [256]int

	fmt.Println(x)
	fmt.Println(x[42])

	for i := 0; i < 256; i++ {
		x[i] = i
	}

	for index, value := range x {
		fmt.Printf("%v - %T and index %v\n", value, value, index)

	}

}
