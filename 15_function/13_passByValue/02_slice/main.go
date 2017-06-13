package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	m[0] = "jain"
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "rajesh"
	fmt.Println(z)
}
