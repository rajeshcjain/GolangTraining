package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main(){

	res,err :=http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil{
		log.Fatalln(err)
	}

	str,err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Fatalln(err)
	}

	words := string(str)
	fmt.Println(words)





}
