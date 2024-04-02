package main

import (
	"fmt"
	"time"
)

// go 1.2x fix
func main() {
	var m = [...]int{1, 2, 3, 4, 5}

	for i, v := range m {
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Println(i, v)
		}()
	}

	time.Sleep(10 * time.Second)
}
