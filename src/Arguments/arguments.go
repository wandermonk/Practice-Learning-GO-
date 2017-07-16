package main

import (
	"fmt"
	"os"
)

//usage : go run arguments.go "hello" "hi"

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
