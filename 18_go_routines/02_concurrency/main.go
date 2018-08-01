package main

import "fmt"


/* Here we have 3 threads which are running...

1). main
2). foo
3). boo

so main thread starts foo and boo thread with go routines
and then it exits and hence foo and boo does not get time to
print any thing...so we need to have some thing to wait.

 */
func main(){
	go foo()
	go boo()
}


func boo(){

	for i:=0;i<100;i++{
		fmt.Println("boo ",i)
	}
}

func foo(){

	for i:=0;i<100;i++{
		fmt.Println("foo ",i)
	}
}