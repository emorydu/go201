package main

import "fmt"

func arrayRangeExpression() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("numbers = ", numbers)

	for i, v := range numbers {
		if i == 0 {
			numbers[1] = 12
			numbers[2] = 13
		}

		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("numbers = ", numbers)
}

func pointerToArrayRangeExpression() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("numbers = ", numbers)

	for i, v := range &numbers {
		if i == 0 {
			numbers[1] = 12
			numbers[2] = 13
		}

		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("numbers = ", numbers)
}

func sliceRangeExpression() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("numbers = ", numbers)

	for i, v := range numbers[:] {
		if i == 0 {
			numbers[1] = 12
			numbers[2] = 13
		}

		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("numbers = ", numbers)
}

func sliceLenChangeRangeExpression() {
	// numbers is slice (-> struct contains len)
	var numbers = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0, 5)

	fmt.Println("numbers = ", numbers)

	for i, v := range numbers { // len immutable
		if i == 0 {
			numbers = append(numbers, 6, 7)
		}

		r = append(r, v)
	}

	fmt.Println("r = ", r)
	fmt.Println("numbers = ", numbers)
}

func main() {
	sliceLenChangeRangeExpression()
}
