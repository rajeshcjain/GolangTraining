package main

import "fmt"

func main(){
	/*
	This is the 2nd shorthand way creating the slice.
	In this we have crated the header but still if we have to use it
	then first use append()
	*/
	student := []string{}
	students := [][]string{}

	//So here if i try access the student[0] then it will give error as
	// we are trying to access index which does not exist as of now.

	//so the write way to access it is using append()
	fmt.Println(student)

	/*
	   When we create the slice using
	   student := []string{}
	   then it creates the array but it does not have any length and capacity..
	   which can not be used...so when we try to use
	   student[0] = "R1"

	   It will give "index out of range" error
	*/


	student = append(student,"Todd")

	fmt.Println(student)

	/*
	  Now lets declare a slice with length and capacity then try to access it.
	  We will see in this case ,we are able to access it.Its because
	  we have declared a slice with length and capacity as 5.So we have already
	  declared and array under the hood with a length of 5 where as in previous case
          shorthand case,we did not have any array under the hood.
 	*/

	stu := make([]string,5)
	stu[0] = "R1"
	fmt.Println(stu)

	fmt.Println(students)
	fmt.Println(student == nil)
	fmt.Println(students == nil)

}
