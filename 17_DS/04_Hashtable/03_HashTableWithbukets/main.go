package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

func main(){

	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")

	if err != nil{
		log.Fatalln(err)
	}

	scanner:= bufio.NewScanner(res.Body)
        defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	//Declare the slice.
	buckets := make([]int,400)

	//Here put the value in the bucket and increment it by 1
	for scanner.Scan(){
		n := HashBucket(scanner.Text())
            	buckets[n]++
	}

	fmt.Println(buckets[65:123])


}


//Here we are returning the integer value of the first character of the string
func HashBucket(str string) int {
      return int(str[0])
}
