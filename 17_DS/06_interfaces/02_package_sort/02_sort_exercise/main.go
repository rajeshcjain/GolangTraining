package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(s)

	//This could have been done in a single line of code.But for better understanding
	//I am doing it in two steps....here i StringSlice is a type which has underlying slice of string
	//So we can convert the slice of string in to stringSlice same as we do for
	// int(value) and string([]byte)
	stringSlice := sort.StringSlice(s)

	/*
		Now as we have converted it into StringSlice,we have lot of functions available with it...
		like

		func (p StringSlice) Len() int
		func (p StringSlice) Less(i, j int) bool
		func (p StringSlice) Search(x string) int
		func (p StringSlice) Sort()
		func (p StringSlice) Swap(i, j int)
	*/

	stringSlice.Sort()
	fmt.Println(stringSlice)

	//This is a 2nd way of doing it as we have already converted it in to
	//string slice we just need to pass it in to sort method.
	sort.Sort(stringSlice)
	fmt.Println(stringSlice)

}
