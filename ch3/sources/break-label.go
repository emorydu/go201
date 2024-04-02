package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan struct{})

	go func() {
	loop:
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("tick")
			case <-exit:
				fmt.Println("exiting...")
				break loop
			}
		}
	}()

	time.Sleep(3 * time.Second)
	exit <- struct{}{}

	time.Sleep(3 * time.Second)
}
