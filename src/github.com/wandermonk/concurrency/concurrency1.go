package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func PrintCounts(label string,count chan int){
	defer wg.Done()
	for {
	val, ok:= <-count
	if !ok {
		fmt.Println("The channel is closed.\n")
		return
	}
	fmt.Printf("The count received is %d from %s\n",val,label)
	if val==10 {
	fmt.Printf("The channel is closed for %s\n",label)
	close(count)
	return
	}
	val++
	count<-val
	}
}

func main() {
	count := make(chan int)
	wg.Add(2)
	go PrintCounts("routine-1",count)
	go PrintCounts("routine-2",count)
	count <- 1
	fmt.Println("Waiting to terminate\n")
	wg.Wait()
	fmt.Println("Terminating the program\n")
}
