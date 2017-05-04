package main

import "fmt"



//IMPORTANT
func main(){


	myFamilyName := "sakshi"

	switch {
	// we can not compare the myFamilyName to 2 as it is of type 2..where as if we use it with len function which accept the string and return the length of it.It is OK
	case len(myFamilyName) == 2 :
		fmt.Println("sakshi and atishay")
		fallthrough
	case myFamilyName == "sakshi" :
		fmt.Println("rajesh")

	default:
		fmt.Println("Leave")
	}
}
