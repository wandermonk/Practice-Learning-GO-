package main

import (
	"fmt"
)

type emptyStruct struct{}

func iter(n int) []emptyStruct {
	return make([]emptyStruct,n)
}

func main() {
	iterable := iter(10)
	for i,v := range iterable {
		fmt.Println("The index is %d, The value is %+V",i,v)
	}
}
