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

		//If you ru the go run -race main.go..you will see that you will get the race condition
		//Here wg is a common variable and we try to modify the wg in both routines..which is a big
		//NO No in Go as in both the routines we are modifying the wg...we are trying to add
		//which is causing the race condition.
		wg.Add(1)
		for i := 0; i < 10; i++ {
			fmt.Println("Putting data in go routine 1 ", i)
			c <- i
		}
		wg.Done()
	}()

	go func() {
		//Now we are trying to add 1 to the wg
		wg.Add(1)
		for i := 0; i < 10; i++ {
			fmt.Println("Putting data in go routine 2 ", i)
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
