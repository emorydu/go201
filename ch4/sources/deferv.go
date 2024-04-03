package main

import "fmt"

func main() {
	fmt.Println("foo1 result:")
	foo1() // 3 2 1 0
	fmt.Println("foo2 result:")
	foo2() // 3 2 1 0
	fmt.Println("foo3 result:")
	foo3() // go 1.20- 4 4 4 4 // go 1.20+ fixed
}

func foo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func foo2() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func foo3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
