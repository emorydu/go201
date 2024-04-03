package main

import "fmt"

func main() {
	foo1()
	foo2()
}

func foo1() {
	nums := []int{1, 2, 3}
	defer func(n []int) { // func([]int{1, 2, 3})
		fmt.Println(n)
	}(nums)

	nums = []int{3, 2, 1}
	_ = nums
}

func foo2() {
	nums := []int{1, 2, 3}
	defer func(p *[]int) { // func(&nums)
		fmt.Println(*p)
	}(&nums)

	nums = []int{3, 2, 1}
	_ = nums
}
