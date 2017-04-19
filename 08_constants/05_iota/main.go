package main

import "fmt"

const (
	_ = iota //0 and discarded with blank identifiers
	B //1
	C  // 2
)

func main(){
	fmt.Println(B)
	fmt.Println(C)
}
