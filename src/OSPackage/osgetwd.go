package main

import (
	"fmt"
	"os"
)

func main() {
	dir, error := os.Getwd()
	if error == nil {
		fmt.Println(dir)
	} else {
		fmt.Println(error)
	}

	out := os.Environ()
	fmt.Println(out)
}
