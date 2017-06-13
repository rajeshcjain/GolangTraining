package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

/*
In this program we are trying to achieve the count of words and the listing the
words in

*/

func main() {

	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")

	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(res.Body)
	scanner.Split(bufio.ScanWords)

	//create a map
	buckets := make(map[string]int, 26)

	for scanner.Scan() {

		str := scanner.Text()
		charAtIndexZero := string(str[0])
		if _, exists := buckets[charAtIndexZero]; exists {
			buckets[charAtIndexZero]++
		} else {
			buckets[charAtIndexZero] = 1
		}
	}

	/*
		Remember this
		range :-

		when used with map...it returns key,value
		when used with slice or array...it returns value,index

	*/

	for k, v := range buckets {
		fmt.Println("starts with ", k, "occurances ", v)
	}
}
