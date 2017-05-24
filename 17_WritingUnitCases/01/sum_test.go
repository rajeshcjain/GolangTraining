package main

import "testing"

func TestSum(t *testing.T){
	total := sum(5,5,)
	if total != 10{
		t.Errorf("expected : %d, but was %d",10,total )
	}
}

