package main

import "fmt"

func main() {
	numbers := []int{11, 12, 13, 14, 15}
	fmt.Println("numbers:", numbers)

	sub1 := numbers[1:3]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(sub1), cap(sub1), sub1)

	sub1 = append(sub1, 24)
	fmt.Println("after append 24, array:", numbers)
	fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(sub1), cap(sub1), sub1)

	sub1 = append(sub1, 25)
	fmt.Println("after append 25, array:", numbers)
	fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(sub1), cap(sub1), sub1)

	sub1 = append(sub1, 26)
	fmt.Println("after append 26, array:", numbers)
	fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(sub1), cap(sub1), sub1)

	sub1[0] = 22
	fmt.Println("after reassign 1st elem of slice, array:", numbers)
	fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(sub1), cap(sub1), sub1)
}
