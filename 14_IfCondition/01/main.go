package main

import (
	"fmt"
)

func main(){

	var i = 0

	fmt.Println("Enyter the value :")
	fmt.Scanf("%d",&i)

	if i < 12{
		fmt.Println("smaller the 12")
	}else {
		fmt.Println("Larger then 12")
	}

}
