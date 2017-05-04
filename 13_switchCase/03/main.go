package main

import "fmt"


//IMPORTANT
func main(){
// fallthrough tells that if the switch option is met and if it has fallthrough written in that switch in end
// then continue to the next case as well. like it is written below.

	switch "sakshi"{
	case "sakshi","atishay" :
		fmt.Println("sakshi and atishay")
		fallthrough
	case "rajesh" :
		fmt.Println("rajesh")

	default:
		fmt.Println("Leave")
	}
}
