package main

import (
	"fmt"
	"sort"
)

func main() {
	studyGrp := []string{"Al", "Zeno", "John", "Jenny"}

	fmt.Println(studyGrp)

	//Converting string in to stringslice
	sliceOfString := sort.StringSlice(studyGrp)
	//Reverse returns the reverse order for data.
	inter := sort.Reverse(sliceOfString)
	//Now sort it
	sort.Sort(inter)
	fmt.Println(sliceOfString)
}
