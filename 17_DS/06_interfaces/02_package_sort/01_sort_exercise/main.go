package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int { return len(p) }

/*
This is a interesting way of swapping the values...
here first p[i] will swap with p[j] and p[j] will be swapping with p[i].Its a position
swap

*/
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i][0] < p[j][0] }

func main() {

	studyGrp := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGrp)
	sort.Sort(studyGrp)
	fmt.Println("After the sort")
	fmt.Println(studyGrp)
}
