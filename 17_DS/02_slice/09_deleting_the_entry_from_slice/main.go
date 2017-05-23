package main

import "fmt"


/*This is important program...it tells how can we delete the entry from slice..


*/

func main(){

	slic1 := []string{"r1","r2","r3","r4"}
	slic2 := []string{"r5","r6","r7","r8"}

	fmt.Println(slic1)
	fmt.Println(slic2)

	//appending one slice with other slice
	slic1 = append(slic1,slic2...)
	fmt.Println(slic1)

	// little lame here...we are deleting the 2nd entry from the slice

	slic1 = append(slic1[:2],slic1[3:]...)
	fmt.Println(slic1)



}

