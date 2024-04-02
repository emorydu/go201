package main

import "fmt"

// go 1.2x fix
func main() {
	var m = [...]int{1, 2, 3, 4, 5}
	for i, v := range m {
		fmt.Printf("i addr: %p, v addr: %p", &i, &v)
		fmt.Println(i, v)
	}
}

/*
var m = [...]int{1, 2, 3, 4, 5}
{
	i, v := 0
	for i, v := range m {
		fmt.Println(i, v)
	}
}
*/
