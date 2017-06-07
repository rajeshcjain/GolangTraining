package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	First string   `json : "first"`
	Last string
	Age int        `json : "age"`
}

func main(){

	//Here are declaring one slice of type byte and the then putting the data init
	bs := []byte(`{"first":"James","Last":"Bond","age":24}`)
	var p1 Person

	//UnMarshalling is used when we get the input from the request and want to produce some
	//meaning full structure out of it..
	err := json.Unmarshal(bs,&p1)

	if err != nil{
		panic(err)
	}

	fmt.Println(p1)
	fmt.Printf("\n%T\n",p1)


}

