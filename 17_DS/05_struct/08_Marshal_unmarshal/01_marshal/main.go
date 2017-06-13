package main

import (
	"encoding/json"
	"fmt"
)

//We are keeping the name of type and field name is capitalized...
//Becoz we can export it
type Person struct {
	First      string
	Last       string
	Age        int
	noExported int //AS this field is not in capital letters so we will not get it as json output
}

/*
Marshal is used to convert struct in to JSON encoding of struct...This is primarily used with string


So when we want convert the struct and send the json response then we marshal the struct.
*/

func main() {

	p1 := Person{First: "R1", Last: "J1", Age: 32, noExported: 007}

	bs, _ := json.Marshal(p1)

	fmt.Println(bs)
	fmt.Printf(string(bs))
	fmt.Printf("\n%T\n", bs)

}
