package main

import "fmt"

func main() {

	myMap := map[int]map[string]string{
		1: map[string]string{
			"A1": "B1",
			"A2": "B2",
		},
		2: map[string]string{
			"A12": "B12",
			"A22": "B22",
		},
		3: map[string]string{
			"A123": "B123",
			"A223": "B223",
		},
	}

	fmt.Println(myMap)

}
