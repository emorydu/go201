package main

import (
	"fmt"
	"time"
)

func main() {
	// recvFromUnbufferedChannel()
	recvFromNilChannel()

}

func recvFromUnbufferedChannel() {
	c := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c { // channel copy -> struct, channel no data, for-range block, wait channel close!
		fmt.Println(v)
	}
}

func recvFromNilChannel() {
	var c chan int // nil channel, deadlock

	for v := range c {
		fmt.Println(v)
	}
}
