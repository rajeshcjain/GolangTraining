package main

import "fmt"

/*
                          IMPORTANT

  This is important,Here we are declaring the slice "transactions" as cap =3 and
  len=3.Now when we call append on that it will append it after 3 as that is the
  existing size.If we change it from make([][]int,3) to make([][]int,0,3) then it will
  better as the current size of the array is ZERO.So when we append it;It will
  append it from ZERO and hence will give the right values

  Otherwise there will be first 3 entries which will blank followed by
  right content...that means content of "transaction" entries.

*/

func main(){
	transactions := make([][]int,3)
	for i := 0; i < 3; i++{
		transaction := make([]int,4)
		for j := 0;j < 4; j++{
			transaction = append(transaction,(j*i))
		}
		transactions = append(transactions,transaction)
	}
	fmt.Println(transactions)
}
