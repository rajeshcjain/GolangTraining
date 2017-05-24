package main

import "fmt"

/*
                                  IMPORTANT
      This program will explain how can we define multi-dimension slice
*/


func main(){

	/*
	  Here we are declaring the 2-dimension slice...It is read as
	  It is a slice of type slice.
	*/
	records := make([][]string,0)

	fmt.Println("**********Declaring the 2-Dimension Slice**********")
	fmt.Println(records)

	student1 := make([]string,4)
	student1 = append(student1,"R1")
	student1 = append(student1,"99%")
	student1 = append(student1,"98%")
	student1 = append(student1,"97%")
	student1 = append(student1,"99%")

	fmt.Println("1st entry of slice")
	fmt.Println(student1)

	student2 := make([]string,4)
	student2 = append(student2,"R2")
	student2 = append(student2,"89%")
	student2 = append(student2,"88%")
	student2 = append(student2,"87%")
	student2 = append(student2,"89%")

	fmt.Println("2nd entry of slice")
	fmt.Println(student2)

	/*So here we are appending the 2-dimensional slice*/
	records = append(records,student1)
	records = append(records,student2)


	fmt.Println("*******************2-Dimensional Slice************************")
	fmt.Println(records)



}
