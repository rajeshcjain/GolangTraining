package main

import "fmt"

func main() {

	myGreetings := map[string]string{
		"R1": "J1",
		"R2": "J2",
		"R3": "J3",
	}

	for index, val := range myGreetings {
		fmt.Println("Index ", index, "Value ", val)
	}
}
