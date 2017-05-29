package main

import "fmt"

func main(){

	var x[256] string

	fmt.Println(x)
	fmt.Println(x[42])

	for i := 0;i < 256; i++{
		x[i] = string(i)
	}

	for _,val := range x{
		fmt.Printf("%v - %T - %v\n", val,val,[]byte(val))
	}

}
