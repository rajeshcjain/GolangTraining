package main

import "fmt"

func main(){
	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
		"J" : "K",
		"L" : "M",
	}

	fmt.Println(myGreeting)
	delete(myGreeting,"J")
	fmt.Println(myGreeting)



}
