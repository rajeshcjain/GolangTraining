package main

import "fmt"

func main(){
	//Assertions are used with interface
	var name interface{} = "Sydney"
	//Here i am doing the assertion as i know name
	//internally of type string
	//remember we can not do type conversion
	//as type conversion happens on actual types
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
