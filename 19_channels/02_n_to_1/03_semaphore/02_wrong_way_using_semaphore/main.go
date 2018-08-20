package main

import "fmt"

func main(){

	c:= make(chan int)
	done := make(chan bool)

	go func() {
		for i:=0;i<10;i++{
			fmt.Println("The value in the 1st routine is ",i)
			c <- i
		}
	}()

	go func() {
		for i:=0;i<10;i++{
			fmt.Println("The value in the 1st routine is ",i)
			c <- i
		}
	}()


	//Now here after 1st/2nd go routine put the value in the channel and wait for rage to pick the value
	//put in channel using the range....but as <-done is part of the main rotine...it will stuck here
	// now value will never be consumed as by range. it will be blocked always.
	<-done
	fmt.Println("The first go routine is done with the putting data in the channel c")
	<-done
	fmt.Println("The Second go routine is done with the putting data in the channel c")
	close(c)

	for n := range c{
		fmt.Println("The output from the channel is ",n)
	}


}
