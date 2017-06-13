package main

import "fmt"

func main() {

	fmt.Println(greet("rajesh", "jain"))
}

// Here we are specifying that in (s string),s will be returned from the
// greet function so as the function know what it has to return
// it just need to specify the return.No need of putting the s with the return.
func greet(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
