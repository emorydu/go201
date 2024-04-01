package main

import "fmt"

func main() {
	var s []int
	m := map[int]int{
		0: 10,
		1: 11,
		2: 12,
	}

	for k, _ := range m {
		s = append(s, k)
	}

	for i := 0; i < len(m); i++ {
		doIteration(s, m)
	}
}

func doIteration(s []int, m map[int]int) {
	fmt.Printf("{ ")
	for _, k := range s {
		v, ok := m[k]
		if !ok {
			continue
		}
		fmt.Printf("[%d, %d] ", k, v)
	}
	fmt.Printf("}\n")
}
