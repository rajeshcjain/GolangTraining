package main

import (
	"sync"
	"fmt"
)

func main(){

	var wg sync.WaitGroup
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
		wg.Wait()
		close(c)
	}()

	for val := range c{
		fmt.Println("extract",val)
	}

}
