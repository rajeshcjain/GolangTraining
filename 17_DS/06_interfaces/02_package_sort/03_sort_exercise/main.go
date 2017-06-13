package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	/*
		Here we will understand the concept of intSlice
	*/

	fmt.Println(n)

	intslice := sort.IntSlice(n)
	sort.Sort(intslice)
	fmt.Println(intslice)

}
