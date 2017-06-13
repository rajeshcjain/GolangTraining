package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First string
	Last  string
	Age   int
}

/*
encode means we are changing the encode it means change in to which humans can not understand and
then send it over stream...For it we need writer to write.
*/

func main() {

	p1 := Person{"James", "Bond", 32}

	json.NewEncoder(os.Stdout).Encode(p1)
}
