package main

import "fmt"


type Contact struct {
	greeting string
	name     string
}


//IMPORTANT
//we are telling to function,we dont know the type of x.So we pass interface.
func SwitcOnType(x interface{}){

	switch x.(type) {
	case int :
		fmt.Println("interger")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("No good type")
	}

}

func main(){

	SwitcOnType(7)
	SwitcOnType("rajesh")
	SwitcOnType(Contact{"rajesh jain here","rajesh"})

}


