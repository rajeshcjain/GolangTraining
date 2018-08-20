package main

import "fmt"

func main() {

	fmt.Println(greet("sakshi", "jain"))
}



func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
