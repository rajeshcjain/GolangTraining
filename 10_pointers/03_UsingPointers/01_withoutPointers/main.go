package main

import "fmt"

func Zero(x int) {
	x = 0
}

func main() {
	x := 10
	fmt.Println("X is ", x)
	Zero(x)
	fmt.Println("X is ", x)
}
