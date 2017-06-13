package main

import "fmt"

func factorial(n int) int {

	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fact := factorial(4)
	fmt.Println(fact)
}
