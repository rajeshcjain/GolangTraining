package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Blank identifiers are used for in the situation where you want to ignore the return values....
// like in this case if the error occurs then we are telling please throw in the bin.
//This sort of code is not acceptable in production env.but good to have otherwise.
func main() {

	res, _ := http.Get("http://www.geekwiseacademy.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println("%s", page)

}
