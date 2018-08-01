package main

import (
	"fmt"
	"sync"
	"time"
)

// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of
// goroutines to wait for. Then each of the goroutines
// runs and calls Done when finished. At the same time,
// Wait can be used to block until all goroutines have finished.

/*
This is a wait group we are defining here

This WaitGroup actually works same as to CountDownLatche in JAVA

*/

var wg sync.WaitGroup

func main(){
	//Here we are adding a delta of 2 as we have 2 thread for which
	//we have to wait.
	wg.Add(2)
	//triggering the foo thread
	go foo()
	//triggering the boo thread
	go boo()
	//Main thread waits here...to finish of these two tasks
	//When the count down reaches from 2 to 0, main thread
	//can pass wg.Wait() function
	wg.Wait()
	fmt.Println("Main thread done here")
}


func boo(){

	for i:=0;i<100;i++{
		fmt.Println("boo ",i)
		//Pausing the thread
		time.Sleep(time.Duration(10*time.Millisecond))
	}
	//Now we tell the Wait group that we are done so reduce the
	//count by one.
	wg.Done()
}

func foo(){

	for i:=0;i<100;i++{
		fmt.Println("foo ",i)
		//Pausing the thread
		time.Sleep(time.Duration(7*time.Millisecond))
	}
	//Now we tell the Wait group that we are done so reduce the
	//count by one.
	wg.Done()
}