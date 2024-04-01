package main

import (
	"fmt"
	"time"
)

func main() {
	m := map[int]int{
		0: 10,
		1: 11,
		2: 12,
	}

	go func() {
		for i := 0; i < 1000; i++ {
			doIteration(m)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			doWrite(m)
		}
	}()

	time.Sleep(5 * time.Second)
}

func doIteration(m map[int]int) {
	for k, v := range m {
		_ = fmt.Sprintf("[%d, %d] ", k, v)
	}
}

func doWrite(m map[int]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}
