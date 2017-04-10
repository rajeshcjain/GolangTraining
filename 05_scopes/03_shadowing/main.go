package main

import "fmt"


//The function has package scope and accessible in package as well.Remember the function name start with small letter;
//so it is not seen outside of the package
func max(x int) int{
	return 42 + x

}

func main(){
        // Now in the below statement it is called shadowing as the max is a variable not function name.
	max := max(7)
	fmt.Println(max)
}
