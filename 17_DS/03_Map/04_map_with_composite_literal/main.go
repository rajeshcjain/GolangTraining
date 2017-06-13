package main

import "fmt"

func main() {

	/*
		Here we are declaring the map using the short hand.
	*/
	myMap := map[string]int{"r1": 5, "r2": 10}
	fmt.Println(myMap)
	fmt.Println(len(myMap))
	fmt.Println(myMap["r1"])

}
