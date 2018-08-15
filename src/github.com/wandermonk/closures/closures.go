package main

import (
	"fmt"
)

func printMessage() func(string) {
	return func(message string) {
		fmt.Println(message)
	}
}

func outer(name string) {
	text := "Modified" + name
	foo := func(){
	fmt.Println(text)
	}
	foo()
}

func returnFunc(name string) func(){ 
	text := "return modified "+name
	foo := func() {
	fmt.Println(text)
	}
       return foo
}

func main() {
	printFunc := printMessage()
	printFunc("hello")

	outer("phani")
	retfunc := returnFunc("mnew func")
	retfunc()

}
