package main

import "fmt"

const (
	A = iota //0
	B        // No need to put iota,It will take the 0 and increment it to 1
	C        // it will increment it to 2
)

const (
	D = iota // it will set the D to 0 again
	E        // 1
	F        // 2
)

func main() {

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
