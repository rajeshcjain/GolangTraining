package main

import "fmt"

func main() {
	fmt.Println(greet())
	fmt.Println(greet("sakshi", "jain"))
}

func greet() string {
	return "rajesh jain"
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
