package main

import (
	"fmt"
)

func main() {

	var s [10]int

	s[0] = 10
	s[1] = 20
	s[2] = 30
	s[4] = 40
	s[5] = 50
	s[6] = 60
	s[7] = 70
	s[8] = 80
	s[9] = 90

	fmt.Println("The size of the array is ", len(s))

	for i, v := range s {
		fmt.Println("the key and value are %d:%d \n", i, v)
	}

	slice1 := []int{1, 2, 3, 4}
	slice1 = append(slice1, 10)

	for i, v := range slice1 {
		fmt.Println("The vkey and value in slice 1 is ::: \n", i, v)
	}

}
