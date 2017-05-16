package main

import "fmt"

// This is a program where defer has not been used to understand what is the
// use of defer.

func hello(){
	fmt.Println("hello")
}

func world(){
	fmt.Println(" world")
}

func main(){
	hello()
	world()

}
