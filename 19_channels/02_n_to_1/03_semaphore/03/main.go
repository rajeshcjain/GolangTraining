package main

import (
	"fmt"
)

/*
Here we are using the semaphore with value as 10...and it is used primarily for closing the
channel c.
*/

func main(){

	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++{
	    //Here we are creating 10 go routines
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println("putting value ",i)
				c <- i
			}
			fmt.Println("setting done as true for ",i)
			done<-true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
			fmt.Println("I am done with value ", i )
		}
		fmt.Println("Closing the channel ",c)
		close(c)
	}()

	for val := range c{
		fmt.Println("The value is ",val)
	}
}
