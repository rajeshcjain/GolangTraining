package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {

	bs := []byte(`{"First":"James","Last":"Bond","Age":23}`)

	//This is a simulation of json coming from somewhere as a stream..So we need to crearte
	//the reader.
	reader := strings.NewReader(string(bs))

	var p1 Person
	err := json.NewDecoder(reader).Decode(&p1)

	if err != nil {
		panic(err)
	}

	fmt.Println(p1)

}
