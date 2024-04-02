package main

import "fmt"

var (
	a = c + b // 5 + 4 = 9
	b = f()   // 4
	c = f()   // 5
	d = 3     // 5
)

func f() int {
	d++
	return d
}

func main() {
	fmt.Println(a, b, c, d)
}
