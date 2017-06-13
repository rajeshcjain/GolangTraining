package main

import "fmt"

func main() {
	myMap := map[string]string{
		"R1": "J1",
		"R2": "J2",
	}

	myMap["R3"] = "J3"

	fmt.Println(myMap)
}
