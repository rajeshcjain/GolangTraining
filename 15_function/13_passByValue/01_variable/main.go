package _1_variable

import "fmt"

func main() {

	x := 42
	fmt.Println(&x)
	fmt.Println(x)

	changeMe(&x)

	fmt.Println(&x)
	fmt.Println(x)

}

func changeMe(z *int) {
	fmt.Println("Start of chnageMe function")
	fmt.Println(z)
	fmt.Println(*z)

	*z = 24

	fmt.Println(z)
	fmt.Println(*z)

	fmt.Println("End of chnageMe function")
}
