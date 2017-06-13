package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["rajesh"] = 44
	fmt.Println(m)
	fmt.Println("Calling changeMe")
	changeMe(m)
	fmt.Println("After Calling changeMe")
	fmt.Println(m)

}

func changeMe(z map[string]int) {
	z["rajesh"] = 22
}
