package main

import (
	"image/color"
	"fmt"
)

type vehical struct{
	Seats int
	MaxSpeed int
	Color string
}

type car struct{
	vehical
	wheels int
	doors int
}

type boat struct{
	vehical
	Length int
}

type plane struct{
	vehical
	Jet bool
}

func main(){

	car1 := car{vehical{4,200,"Black"},4,4}
	cars := []car{car1,car{},car{}}
	planes := []plane{plane{},plane{},plane{}}
	boats := []boat{boat{},boat{},boat{}}


	for key,val := range cars {
		fmt.Println(key, " - ",val )
	}

	for key,val := range planes {
		fmt.Println(key, " - ",val )
	}

	for key,val := range boats {
		fmt.Println(key, " - ",val )
	}



}
