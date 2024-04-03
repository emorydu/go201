package main

import (
	"fmt"
	"strings"
)

const (
	defaultSep = "-"
)

func main() {
	println(concat(defaultSep, 1, 2))
	println(concat(defaultSep, "hello", "world"))
	println(concat(defaultSep, 1, uint32(2), []int{11, 12, 13}, 17, []string{"ED", "E", "D"}, "hacker", 33))
}

func concat(sep string, args ...interface{}) string {
	var r string

	for i, v := range args {
		if i != 0 {
			r += sep
		}
		switch v.(type) {
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			r += fmt.Sprintf("%d", v)
		case string:
			r += fmt.Sprintf("%s", v)
		case []int:
			ints := v.([]int)
			for i, v := range ints {
				if i != 0 {
					r += sep
				}
				r += fmt.Sprintf("%d", v)
			}
		case []string:
			strs := v.([]string)
			r += strings.Join(strs, sep)
		default:
			fmt.Printf("the argument type [%T] is not supported", v)
			return ""
		}
	}

	return r
}
