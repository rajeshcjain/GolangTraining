package main

import "fmt"


//Package level scope
var x int = 42

func main(){

	fmt.Println(x)
	foo();
	//block level scope
	y := 10
	fmt.Println(y)

}

func foo(){
	fmt.Println(x)
}