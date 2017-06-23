package main

import "fmt"

func main(){

	c:=make(chan int)
	go func(){
		for i:=1;i<=10;i++{
			fmt.Println("putting value ",i)
			c<-i
		}
		close(c)
	}()

	for val := range c{
		fmt.Println("extracting",val)
	}
}
