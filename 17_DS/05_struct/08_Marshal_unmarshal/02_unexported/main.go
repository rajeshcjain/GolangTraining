package main

import (
	"fmt"
	"encoding/json"
)


//As all the fields are not exported..i.e they are not in capital letters..so they will not be printed after marshaling
type Person struct{
	first string
	last string
	age int
}

func main(){

	p1 := Person{"james","bond",24}
	fmt.Println(p1)
	bs,_:= json.Marshal(p1)
	fmt.Println(string(bs))
}
