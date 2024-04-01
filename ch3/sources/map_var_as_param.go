package main

import "fmt"

func main() {
	m := map[string]int{
		"ok": 200,
		"ko": -200,
	}

	fmt.Println(m)
	foo(m)
	fmt.Println(m)
}

func foo(m map[string]int) {
	m["ok"] = 0
	m["ko"] = -1
}
