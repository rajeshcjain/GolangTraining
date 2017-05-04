package main

import (
	"math"
	"fmt"
)

func main(){
	var lim float64 = 12
	var enteredValue float64
	fmt.Println("Enter the Value : ")
	fmt.Scan(&enteredValue)

	//First way
	//Here v scope is If block only.
	if v := math.Sqrt(enteredValue); v > 3{
		fmt.Println(v)
		return
	}

	fmt.Println(lim)


	//Second way

	if v := math.Sqrt(enteredValue); v > 5{
		fmt.Println(v)
		return
	}else{
		fmt.Println("I am in the elase part")
		fmt.Println(lim)
	}


}
