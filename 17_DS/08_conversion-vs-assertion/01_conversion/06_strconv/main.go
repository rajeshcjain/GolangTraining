package main

import (
	"strconv"
	"fmt"
)

func main(){

	//converting the string into 45
	num,_ := strconv.Atoi("45")
	fmt.Printf("%T ",num)
	fmt.Println(num)

	str:=strconv.Itoa(45)
	fmt.Printf("%T ",str)
	fmt.Println(str)


}
