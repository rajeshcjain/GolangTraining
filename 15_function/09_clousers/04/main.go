package main

import "fmt"

func wrapper()  func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main(){
	increment := wrapper()
	fmt.Printf("%T \n",increment)
	fmt.Println(increment())
	fmt.Println(increment())
}
