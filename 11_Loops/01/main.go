package main

import "fmt"

func main(){

	//First Way

	sum := 0
	for i := 0;i < 20;i++{
		sum += i
	}
	fmt.Println(sum)

	//Second Way

	sum1 := 0
	for ;sum1 < 1000; sum1 +=20{

	}

	fmt.Println(sum1)

        // Third way

	sum2 := 0
	for sum1 < 10000 {
		sum2 += 20
	}
	fmt.Println(sum2)

}
