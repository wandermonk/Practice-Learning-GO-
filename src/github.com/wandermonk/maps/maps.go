package main

import (
	"fmt"
)

func main() {
	var greetings map[int]string

	greetings = make(map[int]string)

	greetings[1] = "hello"
	greetings[2] = "hi"
	greetings[3] = "good morning"

	for k, v := range greetings {
		fmt.Printf("Key %d : value %s\n", k, v)
	}


	if value,ok := greetings[2];ok {
		fmt.Println(value)
	}else {
		fmt.Println("The key does not exist")
	}
}
