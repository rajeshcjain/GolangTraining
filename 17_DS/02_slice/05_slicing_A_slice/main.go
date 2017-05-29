package main

import "fmt"

func main(){

	x := []string{
		"R1",
		"S1",
		"AJ1",
		"AJLB",
		"AJLS",
		"GJ",
		"RJ",
		"KJ",
	}

	fmt.Println(x)
	fmt.Println(x[2])
	/*
	   It will print the values starting from index 1 and going till index where index 2 is not included..so
	   in this case only index 1 is printed
	*/
	fmt.Print("[1:2]")
	fmt.Println(x[1:2])
	/*
	   IT will start printing from index 0 and it will go till 2 where 2 is not included...so it will print
	   index 0 and 1
	*/
	fmt.Print("[:2]")
	fmt.Println(x[:2])
	/*
	  IT will start from index 5 where 5th index is not included
	  and will print till end.
	*/
	fmt.Print("[5:]")
	fmt.Println(x[5:])
	/*
	  IT will print all the items i.e starting from 0 till max length.
	*/

	fmt.Print("[:]")
	fmt.Println(x[:])

}
