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


func main() {

	sliceOfVehicals := []vehicals{car{}, car{}, car{},plane{}, plane{}, plane{},boat{}, boat{}, boat{}}

	for key,val := range sliceOfVehicals{
		fmt.Println(key , "-", val)
	}
}
