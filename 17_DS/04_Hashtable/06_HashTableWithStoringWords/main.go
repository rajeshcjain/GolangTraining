package main

import (
	"net/http"
	"bufio"
	"fmt"
)

func main(){
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil{
		panic(err)
	}

	scanner := bufio.NewScanner(res.Body)
	scanner.Split(bufio.ScanWords)
	defer res.Body.Close()


	//So here we are making a slice of slice.So here we have a slice with 12 length and
	//we want 12 slices at those 12 locations
	hashTable := make([][]string,12)

	//Then append those 12 locations with slices
	for i := 0; i<12;i++ {
		hashTable = append(hashTable, []string{})
	}

	bs := scanner.Bytes()
	fmt.Println("bs len ", len(bs))


	for scanner.Scan(){
              str := scanner.Text()
		n:= HandleBuket(str)
		hashTable[n] = append(hashTable[n],str)
	}

	for i := 0; i <12 ; i++{
		fmt.Printf("---- %s",hashTable[i])
	}
}

func HandleBuket(word string) int{
	val := int(word[0])
	return val%12
}
