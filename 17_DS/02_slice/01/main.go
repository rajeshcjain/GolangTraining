package main

import "fmt"


//With Array,They are only created statically but slice is a list which can be created dynamically...which means
//that its size can grow and shrink based on requirement like ArrayList in Java
func main(){


	//1). That is how we create the slice in Go
	sliceStr :=[]string{"rajesh","Sakshi","Atishay","AtishayRealBrother","AtishayRealSister"}
	sliceOfint :=[]int{1,2,3,4}

	//2). sliceStr := make([]string,3)
	//This will create an empty slice with cap=3 and len=3
	/*
	    3). sliceType := make([]string,length,capacity)
                sliceStr := make([]string,50,100)

                So here we are creating the slice with name sliceStr which has string type of under lying array
                and the total capacity of array is 100 where as the initial length of the array is 50.So now when we
                put 51th array then the length of the array is changed from 50 to 100;

                That is how the length of the array is increased dynamically;

                 Now suppose that the total initial capacity of the array is 100 and it is full;Now in normal case
                 if we put the 101th element in the array it will throw an error but in the case of slice if we put
                 101th element;It will first create an array of double the size of previous array and then it will copy the
                 content from the old array to new array and then discard the old array and then we can add the new element in
                 to it.
                */

	/*
	  4). sliceStr := new([100]int)[0:50]
	  Its another way of creating the slice.
	*/
	fmt.Printf(" %T \n",sliceStr)
	fmt.Printf(" %T \n",sliceOfint)

	fmt.Println(sliceStr[3])

	/*
	         IMPORTANT
	         Slicing the slice.
	         This means that print the element in the slice starting from 1 and till index 3...3 being not
	         included in the list....so it will print the content on index 1 and 2,not on 3...so "Sakshi","Atishay"
	         will not be printed.
	*/
	fmt.Println(sliceStr[1:3])

	/*
	   We represents the rune of "S" and each rune is represented by unicode and so it will give me the value of
	   unicode.
	*/
	fmt.Println("myString"[2])





}