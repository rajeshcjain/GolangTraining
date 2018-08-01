package main

import "fmt"

type vehicals interface{}

type vehical struct{
	Seats int
	MaxSpeed int
	Color string
}

type car struct {
	vehical
	Door int
	Wheels int
}

type plane struct{
	vehical
	Jet bool
}

type boat struct{
	vehical
	Length int
}


/*

Now some important point here...if i replace it with

func printValues(p []interface{}){
	for key,val := range p{
		fmt.Println(key," - ",val)
	}
}

then this will not work as it will be slice of empty interface
but when we change it to the below declaration...it basically
tells the compiler to send me the values one by


*/
func printValues(p ...interface{}){
	for key,val := range p{
		fmt.Println(key," - ",val)
	}
}

func main(){

	//Uncomment the code..it will give error.

	cars := []car{car{}, car{}, car{}}
	planes := []plane{plane{}, plane{}, plane{}}
	boats := []boat{boat{}, boat{}, boat{}}

	printValues(cars)
	printValues(planes)
	printValues(boats)

}
