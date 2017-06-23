package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){
	c := make(chan int)

	wg.Add(2)

	go func(){
/*
This is how channel are used.We add ZERO to channel and then
it stops the thread.Then 2nd thread starts executing and it will
fetch the value from c and print it.Then the 2nd thread stops executing
and 1st thread fetch the 2nd value and store it in channel again and
then 1st thread stops executing and 2nd thread starts executing and
so on...
*/
		for i:=0;i<10;i++{
			c<-i
		}
		wg.Done()
	}()

	go func(){
		for i:=0;i<10;i++{
			fmt.Println(<-c)
		}
		wg.Done()
	}()
	wg.Wait()
}
