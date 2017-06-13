package main

import "fmt"

func main() {

	myMap := make(map[string]int)
	myMap["k1"] = 12
	myMap["k2"] = 22

	/*
		This is important as myMap[key] returns two values,
		first being value and second being whether that is present or not
	*/
	val, present := myMap["k1"]
	fmt.Println("value ", val, "present ", present)

	delete(myMap, "k1")

	val, present = myMap["k1"]
	fmt.Println("value ", val, "present ", present)

}
