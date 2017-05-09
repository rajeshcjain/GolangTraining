package main

import "fmt"

func main(){

	fmt.Println(greet("rajesh","jain"))
	var str,_ = greet("sakshi","jain")
	fmt.Println(str)

}

// so in golang we can return multiple values.In the function declaration we are specifying
// that it can return two strings with the (string,string)
// and in return statement we are returning the two string which are comma seperated.

func greet(fname,lname string) (string,string){
	return fmt.Sprint(fname,lname),fmt.Sprint(lname,fname)
}