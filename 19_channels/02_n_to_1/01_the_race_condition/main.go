package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)

	//Declaring a variable of type WaitGroup
	var wg sync.WaitGroup

	go func() {
		//Now we are trying to add 1 to the wg
		//These 2 separate threads we are modifying
		//Which is a No No in GoLang
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		//Now we are trying to add 1 to the wg
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
