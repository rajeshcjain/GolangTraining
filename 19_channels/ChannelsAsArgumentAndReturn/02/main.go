package main

import "fmt"

/*
Below program is a fix for 01 program where we encountered the DEAD LOCK.

*/

func main(){
	c := incrementor()
	cSum := puller(c)

	//Here channel cSum function will be ranged and the output will be printed.
	for n := range cSum{
		fmt.Println(n)
	}
}

func puller(c chan int) chan int{
	out := make(chan int)
	var sum int

	go func() {
		for val := range c{
			sum += val
		}
		out<-sum
		close(out)
	}()
	return out
}

func incrementor() chan int {
	c := make(chan int)

	//Here we created the go routine which will run independently of incrementor function.

	//incrementor function will accept the spawn this go routine and then it will return it
	//that way puller() function will keep on executing.
	//Now here go routine will put the value to channel c and then will wait...
	//puller function where we are doing the same thing except that it accepts channel c as a
	//input;starts fetching the value and these 2 go routines will go back and fourth.
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
