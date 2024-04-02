package main

import "fmt"

var (
	a    = c   // a = 5
	b, c = f() // b = 4 c = 5
	d    = 3   // d = 4
)

func f() (int, int) {
	d++
	return d, d + 1
}

func main() {
	fmt.Println(a, b, c, d) // 5 4 5 4
}
