package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case (getASlice())[0] = <-getAReadOnlyChannel():
		fmt.Println("recv something from a readonly channel")
	case getAWriteOnlyChannel() <- getANumToChannel():
		fmt.Println("send something to a writeonly channel")
	}
}

func getAReadOnlyChannel() <-chan int {
	fmt.Println("invoke getAReadOnlyChannel")
	c := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		c <- 1
	}()

	return c
}

func getASlice() *[5]int {
	fmt.Println("invoke getASlice")
	var a [5]int

	return &a
}

func getAWriteOnlyChannel() chan<- int {
	fmt.Println("invoke getAWriteOnlyChannel")
	return make(chan int)
}

func getANumToChannel() int {
	fmt.Println("invoke getANumToChannel")
	return 2
}

// output
// invoke getAReadOnlyChannel
// invoke getAWriteOnlyChannel
// invoke getANumToChannel

// invoke getASlice
// recv something from a readonly channel
