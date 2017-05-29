package main

import (
	"fmt"
)

func main(){

	slic1 := []string{"r1","r2","r3","r4"}
	slic2 := []string{"r5","r6","r7","r8"}

	fmt.Println(slic1)
	fmt.Println(slic2)


	/*
	 Here please mind the slic2...we have to specify it as "slic2..." as
	 func append(slice []Type, elems ...Type) []Type

	 we pass the slice as list of elements in the append function.

	*/
	slic1 = append(slic1,slic2...)
	fmt.Println(slic1)


}
