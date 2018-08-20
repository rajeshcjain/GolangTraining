package main

import "fmt"

func main() {
	//here we are specifying the slice which is of type []float64
	//if we pass the data directly to the avg...then the type we are
	//passing is of type []float64 where as avg accept list of float64.
	//So here we are declaring data which is one item which has
	//bunch of items listed in it.
	data := []float64{21, 23, 44, 88, 9, 13}

	//so to pass one item at a time from slice; we add "." behind the data
	//which tells the avg that pass one item at a time to
	//avg function.
	n := avg(data...)
	fmt.Println(n)

	//This is the case when our function accepts the []float64 that means slice.
	fmt.Println(avg1(data))
}

//Here we are specifying that the avg will accept the unlimited float64 numbers
//so it will be accepting the 1st float64 number then 2nd then 3rd
//If we want to avoid the all circus we could have declared the function as
// func avg(sf []float64) float64 so now in the main function
//we could have done  "n:=avg(data)"...then it will be
// passing []float64 type.
func avg(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)

	var total float64
	for _, v := range sf {
		total += v
	}
	return total
}

//here is the function which accepts one item and that is of type slice of float64
// which we write as []float64
func avg1(sf []float64) float64{
	fmt.Println(sf)
	fmt.Printf("%T \n",sf)

	var total float64
	for _,v := range sf{
		total += v
	}
	return total

}
