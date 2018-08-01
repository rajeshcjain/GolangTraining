package main

import (
	"fmt"
	"sync"
	"runtime"
	"time"
)

func boo(){

	for i:=0;i<100;i++{
		fmt.Println("boo ",i)
		time.Sleep(5*time.Millisecond)
	}
	wg.Done()
}

func boo1(){

	for i:=0;i<100;i++{
		fmt.Println("boo1 ",i)
		time.Sleep(5*time.Millisecond)

	}
	wg.Done()
}

func boo2(){

	for i:=0;i<100;i++{
		fmt.Println("boo2 ",i)
		time.Sleep(5*time.Millisecond)
	}
	wg.Done()
}

func boo3(){

	for i:=0;i<100;i++{
		fmt.Println("boo3 ",i)
		time.Sleep(5*time.Millisecond)
	}
	wg.Done()
}

func boo4(){

	for i:=0;i<100;i++{
		fmt.Println("boo4 ",i)
		time.Sleep(5*time.Millisecond)
	}
	wg.Done()
}

func boo5(){

	for i:=0;i<100;i++{
		fmt.Println("boo5 ",i)
		time.Sleep(5*time.Millisecond)
	}
	wg.Done()
}

func boo6(){

	for i:=0;i<100;i++{
		fmt.Println("boo6 ",i)
		time.Sleep(5*time.Millisecond)
	}
	wg.Done()
}

func foo(){

	for i:=0;i<100;i++{
		fmt.Println("foo ",i)
		time.Sleep(5*time.Millisecond)
	}
	wg.Done()
}

var wg sync.WaitGroup

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main(){

	wg.Add(8)
	go foo()
	go boo()
	go boo1()
	go boo2()
	go boo3()
	go boo4()
	go boo5()
	go boo6()

	wg.Wait()
}
