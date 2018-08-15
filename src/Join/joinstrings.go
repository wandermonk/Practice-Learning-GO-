package main

import (
	"fmt"
	"strings"
)

func main() {
	var t = "hello"
	var s = ""
	s += strings.Join(strings.Split(t, ""), "\t")
	fmt.Println(s)
}
