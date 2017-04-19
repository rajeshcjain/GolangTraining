package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main(){
	res, err := http.Get("http://www.geekwiseacademy.com/")
        if err != nil {
		log.Fatal(err)
	}

	page,err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("%s",page)
}
