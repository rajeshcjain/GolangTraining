package main

import "fmt"

func main() {

	sliceStr := make([]int, 0, 3)

	fmt.Printf("%T\n", sliceStr)
	fmt.Println(sliceStr)

	fmt.Println("Cap ", cap(sliceStr))
	fmt.Println("Len ", len(sliceStr))
	/*************************Appending it now*****************/

	//So here we are appending far ahead of the initial capacity of slice which was 3
	for i := 0; i < 80; i++ {
		sliceStr = append(sliceStr, i)
		fmt.Println("Len : ", len(sliceStr), "cap : ", cap(sliceStr))
	}

}
