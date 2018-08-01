package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

var count int

func main(){
	wg.Add(2)
	go increment("foo:")
	go increment("boo:")
	wg.Wait()
	fmt.Println("Main is finished")
}

func increment(str string){
	x := 0
	for i:=0;i<20;i++{
		x++
		count += x
		fmt.Println(str," count:",count," x:",x)
	}
	wg.Done()
}
