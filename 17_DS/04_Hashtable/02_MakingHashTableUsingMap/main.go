package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil {
		log.Fatalln(err)
	}

	/*
		This is bufio - that is buffer input output

		Here we are creating a scanner whose job is to scan the body.
		After getting the scanner then split the words using the ScanWords
		function which will split it.
	*/
	reader := bufio.NewScanner(res.Body)
	reader.Split(bufio.ScanWords)

	//Created a map
	myMap := map[string]string{}

	// Scan advances the Scanner to the next token, which will then be
	// available through the Bytes or Text method. It returns false when the
	// scan stops, either by reaching the end of the input or an error.
	// After Scan returns false, the Err method will return any error that
	// occurred during scanning, except that if it was io.EOF, Err
	// will return nil.
	for reader.Scan() {
		myMap[reader.Text()] = ""
	}

	//This is how the error handling is done here.
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input :", err)
	}

	fmt.Println(myMap)

}
