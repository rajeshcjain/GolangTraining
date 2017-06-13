package main

import "fmt"

type foo int

//Go is a statically typed..that means ...you can not leave the type conversion to interpreter as in
//case of java script.You have to explicitly convert it.
func main() {

	var myAge foo
	myAge = 44
	fmt.Println(myAge)
	//So here when we check the the type of myAge..then it will return us foo rather than int
	fmt.Printf("%T %v", myAge, myAge)

	var myAge1 int
	myAge1 = 44
	fmt.Printf("\n%T %v", myAge1, myAge1)

	// But if we try to add these myAge and myAge1 then it will give error as these 2 are
	//different types
	//fmt.Println("sum", myAge1 + myAge)

	//But if i convert the myAge to int then it will work as we know that internally foo is of type
	//int
	fmt.Println("\nsum", int(myAge)+myAge1)

}
