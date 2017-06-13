package main

import "fmt"

/*
                            IMPORTANT
   It is important to  understand that here,We are creating the slice student here
   but this slice is empty.Slice has three things in Header.
   1). A pointer to under lying array
   2). length
   3). capacity

   so when we declare the slice with the below way..then we are creating the slice
   which is empty.So when we declare the slice using it then we are creating the
   slice which is nil.

   to append it we have to append it using the append function.

*/
func main() {

	var student []string
	var students [][]string

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	fmt.Println(students == nil)
}
