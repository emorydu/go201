package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3
	println(sum(a, b, c))
	nums := []int{1, 2, 3}
	println(sum(nums...))
	// println(sum(a, b, c, num...))	// compile error

	// s := []string{"ED", "E", "D"}
	// dump(s...)

	s := []interface{}{"ED", "E", "D"}
	dump(s...)

	var buf []byte
	buf = append(buf, "hello, world:)"...)
	fmt.Println(string(buf))

	// foo("hello, world:)"...)
}

func foo(b ...byte) {
	fmt.Println(string(b))
}

func sum(args ...int) int {
	var t int

	for _, v := range args {
		t += v
	}

	return t
}

func dump(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}
