package main

import "fmt"

func f() int {
	fmt.Println("calling f")
	return 1
}

func g(a, b, c int) int {
	fmt.Println("calling g")
	return 2
}

func h() int {
	fmt.Println("calling h")
	return 3
}

func i() int {
	fmt.Println("calling i")
	return 1
}

func j() int {
	fmt.Println("calling j")
	return 1
}

func k() bool {
	fmt.Println("calling k")
	return true
}

func main() {
	var y = []int{11, 12, 13}
	var x = []int{21, 22, 23}

	var c chan int = make(chan int)
	go func() {
		c <- 1
	}()

	y[f()], _ = g(h(), i()+x[j()], <-c), k()
	// f()	calling f 1	// y[1], _ = g(h(), i() + x[j()], <-c), k()
	// h()	calling h 3 // y[1], _ = g(3, i() + x[j()], <-c), k()
	// i()	calling i 1 // y[1], _ = g(3, 1 + x[j()], <-c), k()
	// j()	calling j 1 // y[1], _ = g(3, 22, <-c), k()
	// <-c 1	y[1], _ = g(3, 22, 1), k()
	// g()	calling g 2 // y[1], _ = 1, k()
	// k() 	calling k true
	// y[1] = 1

	// output:
	// calling f
	// calling h
	// calling i
	// calling j
	// calling g
	// calling k
}
