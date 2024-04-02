package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for i, v := range numbers {
		go func(i, v int) {
			time.Sleep(3 * time.Second)
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(10 * time.Second)
}
