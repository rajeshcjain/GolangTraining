package main

import "fmt"

// rune is a character from any language or its a alias of int32.
func main() {

	// here basically we have converted the Hello in to byte stream which represents the byte for each character.
	// Now we get out put as [72 101 108 108 111] where 72 represents the "H" in UTF-8 table, 101 represents "e" and so on.
	fmt.Println([]byte("Hello"))

	for i := 5000; i < 5100; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}

}
