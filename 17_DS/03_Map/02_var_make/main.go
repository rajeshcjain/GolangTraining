package main

import "fmt"

func main(){

	/*
	 This is how you should create a map...using make
	*/

	myMap := make(map[string]string)

	fmt.Println(myMap)
	myMap["R1"] = "J1"
	myMap["S1"] = "J1"

	fmt.Println(myMap)

}
