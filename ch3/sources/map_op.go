package main

import "fmt"

func main() {
	m := make(map[string]string)

	// insert
	m["hello"] = "world" // not exists, append
	m["olleh"] = "dlrow"
	m["emory"] = "du"

	fmt.Println(m)

	m["olleh"] = "world" // exists overwirte
	fmt.Println(m)

	// elements length
	fmt.Println("m len:", len(m))
	m["append"] = "append"
	fmt.Println("m len:", len(m))

	// query & read
	_, ok := m["hello"] // comma ok
	if !ok {
		// key not exists
	}

	v := m["hello"]
	fmt.Println(v)

	// delete
	delete(m, "olleh")
	delete(m, "any") // not exists key, no panic
	fmt.Println(m)
}
