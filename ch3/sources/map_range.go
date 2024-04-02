package main

import "fmt"

func main() {
	m := map[string]int{
		"emorydu": 24,
		"smith":   32,
		"ED":      25,
	}

	cnt := 0

	for k, v := range m {
		if cnt == 0 {
			delete(m, "ED")
			m["created"] = 33
		}
		cnt++
		fmt.Println(k, v)
	}

	fmt.Println("counter is:", cnt)
}
