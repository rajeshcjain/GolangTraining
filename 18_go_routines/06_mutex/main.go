package main

import (
	"sync"
	"fmt"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

var count int

func main(){
	wg.Add(2)
	go increment("foo:")
	go increment("boo:")
	wg.Wait()
	fmt.Println("Main is finished")
}

func increment(str string){
	for i:=0;i<20;i++{
		time.Sleep(time.Duration(5*time.Millisecond))
		mutex.Lock()
		x := count
		x++
		count = x
		fmt.Println(str," count:",count," x:",x)
		mutex.Unlock()

	}
	wg.Done()
}
