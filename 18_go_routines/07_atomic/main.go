package main

import (
	"sync"
	"fmt"
	"time"
	"sync/atomic"
)

var wg sync.WaitGroup


var count int64

func main(){
	wg.Add(2)
	go increment("foo:")
	go increment("boo:")
	wg.Wait()
	fmt.Println("Main is finished :: final count ",count)
}

func increment(str string){
	for i:=0;i<20;i++{
		time.Sleep(time.Duration(5*time.Millisecond))
		atomic.AddInt64(&count,1)
		fmt.Println(str," count:",atomic.LoadInt64(&count))
	}
	wg.Done()
}

//func incrementor(s string) {
//	for i := 0; i < 20; i++ {
//		time.Sleep(time.Duration(5) * time.Millisecond)
//		atomic.AddInt64(&count, 1)
//		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&count)) // access without race
//	}
//	wg.Done()
//}
