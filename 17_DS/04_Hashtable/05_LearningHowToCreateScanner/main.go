package main

import (
	"bufio"
	"strings"
	"fmt"
)

func main(){

	var str = "This is Rajesh Jain.I am a good developer and love to learn new technological stack." +
		"Will you like to work with me."


	/*
	So here something to learn here.We can create a Reader interface and can pass the constant value
	to Scanner and then can use it with Scanner which accepts Reader.
	*/
	scan := bufio.NewScanner(strings.NewReader(str))

	scan.Split(bufio.ScanWords)
	for scan.Scan(){
		fmt.Println(scan.Text())
	}
}
