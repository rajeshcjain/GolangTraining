package main

import "fmt"

func Zero(x *int) {
	*x = 0
}

func main() {
	x := 10
	fmt.Println("The value of x is ", x)
	Zero(&x)
	fmt.Println("The value of x is ", x)

}
