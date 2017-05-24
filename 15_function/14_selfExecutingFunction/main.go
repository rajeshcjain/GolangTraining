package main

import "fmt"


//Below is the anonymous function....it does not need any name and as we are putting the
// "()" after the function body ...it will execute there itself.

func main(){
	func() {
		fmt.Println("This is a anonymous function")
	}()
}
