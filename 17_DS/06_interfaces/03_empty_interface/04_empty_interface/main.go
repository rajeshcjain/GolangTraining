package main

import "fmt"

type animal struct{
	sound string
}

type cat struct {
	animal
	friendly bool
}

type dog struct{
	animal
	friendly bool
}


/*
Now this works here as every struct implements the empty interface
but if this would have been type slice then it is not same as
empty interface...it is a type of slice of empty interface..
which is not same as empty interface

*/
func printVal(p interface{}){
	fmt.Println(p)
}

func main(){

	d := dog{animal{"bark"},true}
	c := cat{animal{"meow"},true}
	printVal(d)
	printVal(c)

}

