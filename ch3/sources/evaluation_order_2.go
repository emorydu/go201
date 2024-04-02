package main

import "fmt"

var (
	a = c + b // c = 6 + 4 = 10
	b = f()   // b = 4
	_ = f()
	c = f() // c = 6
	d = 3   // d = 6
)

func f() int {
	d++
	return d
}

func main() {
	fmt.Println(a, b, c, d)
}
