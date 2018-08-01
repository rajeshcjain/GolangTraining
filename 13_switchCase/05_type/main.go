package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

//IMPORTANT
//we are telling to function,we don't know the type of x.So we pass interface.
func SwitcOnType(x interface{}) {
	// so here type is one of the key word of the language which tells hey..give me the
	// type of x. and if it tunr out to be Contact then run the case Contact.
	switch x.(type) { // here x.(type) is not conversion..it is assertion..we are telling that "x is of this type"
	// so if x is turn out to be int then the below case.
	case int:
		fmt.Println("interger")
	// if type is string then the below case.
	case string:
		fmt.Println("string")
	   // if type is of Contact .. then below is the type.
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("No good type")
	}
}

func main() {

	SwitcOnType(7)
	SwitcOnType("rajesh")
	SwitcOnType(Contact{"rajesh jain here", "rajesh"})

}
