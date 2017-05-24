package main

import "fmt"


/*
Map is the key-Value list.It is used to get the implementation which can
be dictionary type.Important point to note that is,Once you place the value in the
Map and if you print the all values then it does not guarantee that the
order will be same.It can change.It uses Hash table below for implementation.

Other then that an uninitialized map is nil.It does not point to anything...just
to nil.




*/

func main(){

	/*
	It is short hand of creating the map....and again it is nil underlying.
	This is not the way you should declare the map as there are no append
	function available to update the map.So preferably use other ways.
	*/
	var myMap map[string]string
	fmt.Println(myMap)
	fmt.Println(myMap == nil)

	/*
	Here if we assign the it to value
	myMap["R1"] = "J1"
	myMap["S1"] = "J1"

	*/

}
