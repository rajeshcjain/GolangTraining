package main

import "fmt"

func wrapper() func() int{
	x := 0
	return func() int{
		y := 30
		fmt.Println(y)
		x++
		return x
	}
}

func main(){
	increament := wrapper()

	secFunc := func() func() int {
		y := 10
		return func() int {
			y++
			return y
		}
	}

	fmt.Println(increament())
	fmt.Println(increament())
	fmt.Println(secFunc())
}
