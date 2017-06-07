package main

import (
	"encoding/json"
	"fmt"
)


/*
Tags are important to learn as they are used widely....here we are tagging the fields and Last is being
tagged as Last with "-"..which tells that This will not be included in the output json.

*/
type Person struct{
	First string  `json:"first"`
	Last string  `json:"-"`
	Age int       `json:"age"`
}

func main(){
	p1 := Person{"james","bond",24}
	bs,_:= json.Marshal(p1)
	fmt.Println(string(bs))

}

