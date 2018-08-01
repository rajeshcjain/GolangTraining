package main

import "fmt"

func main(){
	var name interface{} = 6

	str, ok := name.(int)

	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
