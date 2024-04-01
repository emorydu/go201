package main

import "fmt"

func main() {
	m := map[int]int{
		0: 10,
		1: 11,
		2: 12,
	}

	for i := 0; i < 3; i++ {
		doIteration(m)
	}

}

func doIteration(m map[int]int) {
	fmt.Printf("{ ")
	for k, v := range m {
		fmt.Printf("[%d, %d] ", k, v)
	}
	fmt.Printf("}\n")
}
