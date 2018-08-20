package main

import "fmt"

/*

This is program will cause a dead lock....The reason i am explaining below.
When we run the main function ... incrementor() will be called...and the incrementor()
has a channel which is putting some value in the channel..but when we are putting value in the
channels..there is nothing to accept that value...i accept that we have a range in the puller()
function but that will never be called as the code is already stuck in the incrementor().

So the program get a DEAD LOCK.....

*/

func main(){
	c := incrementor()
	cSum := puller(c)

	for n := range cSum{
		fmt.Println(n)
	}

}


func puller(c chan int) chan int{
	out := make(chan int)
	var sum int

	for val := range c{
		sum += val
	}
	out<-sum
	close(out)
	return out
}

func incrementor() chan int {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	return c
}


