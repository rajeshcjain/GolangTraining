package main

import "fmt"

func main() {

	/*here we are creating the map and it creates the underlying structure
	and then you can perform the operations straight away on the maps*/
	myMap := map[string]int{}

	fmt.Println(myMap)
	fmt.Println(myMap == nil)

	myMap["R1"] = 10
	myMap["R2"] = 20

	fmt.Println(myMap)
}
