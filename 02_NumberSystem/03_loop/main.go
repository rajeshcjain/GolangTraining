package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("decimal %d, hexadecimal %#X binary %b UTF 8 %q \n", i, i, i, i)
	}
}



