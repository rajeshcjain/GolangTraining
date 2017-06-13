package main

import "fmt"

func main() {

	switch "rajesh" {
	case "sakshi":
		fmt.Println("sakshi")

	case "rajesh":
		fmt.Println("rajesh")
		fallthrough
	case "atishay":
		fmt.Println("atishay")
		fallthrough
	default:
		fmt.Println("Leave")
	}

}
