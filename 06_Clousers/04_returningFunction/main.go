package main

import "fmt"

/*
So here when the wrapper function which returns the function

and when we assign the returned function using

increament := wrapper()

then increment is pointing to function returned from wrapper()

 now when we call increment() then it will has a access to "x" as that has a higher scope to
the returned function. so the value of the x will be persisted across the 2 calls of increment()

as it happens in case of clousers.

but had this x being declared in inside the function return from wrapper then it will be
created in for each function call.

*/
func wrapper() func() int {
	x := 0
	return func() int {
		y := 30
		fmt.Println(y)
		x++
		return x
	}
}

func main() {
	increament := wrapper()

	secFunc := func() func() int {
		y := 10
		return func() int {
			y++
			return y
		}
	}

	fmt.Println(increament())
	fmt.Println(increament())

	//In this call we are actually calling the function which returns a function
	//so this will print the address of that function as the value secFunc returns
	// is function so prints address of that.
	fmt.Println(secFunc())

	//Where as another () after secFunc() calls that function which is return by
	//secFunc() so it will print 11.
	fmt.Println(secFunc()())
	}
