package main

import "fmt"

func main(){
	c := make(chan int)

	//Here done which is unbuffered channel acting as a semaphore.
	done := make(chan bool)

	go func(){
//So in one channel it will keep on putting the values in the channel c
		for i := 0 ; i < 10; i++{
			fmt.Println("The value put in the 1st go routine in the channel c ",i)
			c<-i
		}
		//when it will be done doing it...we set the semaphore to true
		done<-true
	}()

	go func(){
		for i := 0 ; i < 10; i++{
			fmt.Println("The value put in the 2nd go routine in the channel c ",i)
			c<-i
		}
		done<-true
	}()


	//Remember one thing here that we can not move the piece of code out side in the
	//main function like mentioned below
	/*

	<-done
	fmt.Println("The first go routine is done with the putting data in the channel c")
	<-done
	fmt.Println("The Second go routine is done with the putting data in the channel c")
	close(c)

	This is because when we put it in the main routine then if you go by the flow...
	First go routine will put zero in the channel c and then will wait for range clause
	to extract from it...and it will be keep on waiting on <-done....and it will never reach to
	range clause....so it will be stuck for ever...so we need to keep the <-done in a separate
	routine...so that main be a separate routine and it has a access of range clause which
	will absorb the values put in the channel c.

	*/
	go func(){
		//Semaphore will be blocked here till done is true.
		//We are using 2 <-done here as when the first anonymous function puts true in the
		// done channel then the value here will be extracted from the channel and after that
		//there will be nothing to extract from it...then 2nd channel will put true in channel
		// and then that too will be extracted from it...making a way for closing the channel.
		<-done
		fmt.Println("The first go routine is done with the putting data in the channel c")
		<-done
		fmt.Println("The Second go routine is done with the putting data in the channel c")
		close(c)
	}()

	for n := range c{
		fmt.Println("The output from the channel is ",n)
	}

	fmt.Println("I am exiting..")

}
