package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["H"] = "J"
	fmt.Println(myGreeting)
	myGreeting["H"] = "K"
	fmt.Println(myGreeting)
}
