package main

import "fmt"

func main(){

	fmt.Println("These are prefered way of creating the variables....please avoid other way of creating them\n ")

	 a := 12
	 b := "rajesh"
	 c := false
	 d := 12.9


	fmt.Println("===============Printing the default value of variables using %v=====================")

	fmt.Printf("var a = %v \n", a)
	fmt.Printf("var b = %v \n", b)
	fmt.Printf("var c = %v \n", c)
	fmt.Printf("var d = %v \n", d)

	fmt.Println("===============Printing the Go-Syntax of variables using %T=====================")

	fmt.Printf("var a = %T \n", a)
	fmt.Printf("var b = %T \n", b)
	fmt.Printf("var c = %T \n", c)
	fmt.Printf("var d = %T \n", d)


	fmt.Println("=================How to do Zero Value creation of variables which is = to creating variables with initilizing it with default value===============")

	var e int16
	var f float64
	var g string
	var h bool

	fmt.Printf("var e = %v  and type : %T \n", e,e)
	fmt.Printf("var f = %v and type : %T \n", f,f)
	fmt.Printf("var g = %v and type : %T \n", g,g)
	fmt.Printf("var h = %v and type : %T \n", h,h)


}
