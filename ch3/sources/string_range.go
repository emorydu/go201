package main

import "fmt"

func main() {
	s := "中国"
	for i, v := range s {
		fmt.Println(i, string(v))
	}
}
