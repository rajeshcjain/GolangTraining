package main

import (
	"sync"
	"fmt"
)

func main(){

	var wg sync.WaitGroup
	//So here is the solution of the race condition which happen in the last program.
	//we mentioned the wg.add(2)...which is a common space and in the main thread.
	//but we still can call wg.done() in the different go routines.
	// which is acceptable.
	wg.Add(2)

	c := make(chan int)

	go func() {
		for i:=0;i<10;i++{
			c<-i
			fmt.Println("First ",i)
		}
		wg.Done()
	}()

	go func() {
		for i:=0;i<10;i++{
			c<-i
			fmt.Println("Second ",i)
		}
		wg.Done()
	}()

	go func(){
		fmt.Println("This is the place we have declared the wait of waiting group.")
		wg.Wait()
		close(c)
	}()

	for val := range c{
		fmt.Println("extract",val)
	}

}
