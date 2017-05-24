package main

import "fmt"

func main(){
	mySlice := []int{10}

	fmt.Println(mySlice)
	mySlice[0]++
	fmt.Println(mySlice)
}
