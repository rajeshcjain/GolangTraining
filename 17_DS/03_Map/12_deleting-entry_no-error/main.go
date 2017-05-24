package main

import "fmt"

func main(){
	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)
	/*
	So as key does not exists;it does nothing.
	*/

	delete(myGreeting,7)
	fmt.Println(myGreeting)



}
