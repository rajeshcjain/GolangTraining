package main

import "fmt"

func main() {

	var x [60]byte
	fmt.Println(x)
	fmt.Println(x[42])
	for i := 0; i < 60; i++ {
		x[i] = byte(i)
	}
	fmt.Println(x)

	for index, val := range x {
		fmt.Printf("%v - %T - %v", val, val, index)
	}
}
