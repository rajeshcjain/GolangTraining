package main

import "fmt"

func main() {

	/*
		So here we have declared the map of slice..
		map[int][]string --> here []string is a slice.
	*/
	myMap := map[int][]string{
		1: []string{"A1",
			"A2",
		},
		2: []string{"A1",
			"A2",
		},
		3: []string{"A1",
			"A2"},
	}

	fmt.Println(myMap)
}
