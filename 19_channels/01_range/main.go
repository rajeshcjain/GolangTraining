package main

import "fmt"

func main(){

	//Here we are creating the a channel with a name c..here as we are not passing the the size
	//so here channel is unbuffered.
	c:=make(chan int)
	go func(){
		for i:=1;i<=10;i++{
			fmt.Println("putting value ",i)
			c<-i
		}
		//When we close the channel we can not write into it after that...
		close(c)
	}()

	//So here when we put 1 in the for loop above then loop will be put in blocking condition.
	//as channel is waiting for that value to be extracted from the range clause below.
	// then when that is read from the channel...go routine will put another value in the
	//channel c and then again it will be blocked...then again range will be reading it.
	for val := range c{
		fmt.Println("extracting",val)
	}
}
