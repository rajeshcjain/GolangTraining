package main

import "fmt"

//When we put the character in single quotes i.e 'a' then it represents the character and hence rune when we change it to double quotes then it means it will be string so "a" is a string.

func main(){

	fmt.Println([]byte("Hello"))

	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T ",foo)
	fmt.Printf("%T",rune(foo))

}
