package main

import (
	"fmt"
	"reflect"
)

type T struct{}

func (T) M1()  {}
func (*T) M2() {}

type I interface {
	M1()
	M2()
}

type T1 T
type I1 I

func main() {
	var t T
	var pt *T
	var t1 T1
	var pt1 *T1

	DumpMethodSet(&t)
	DumpMethodSet(&t1)

	DumpMethodSet(&pt)
	DumpMethodSet(&pt1)

	DumpMethodSet((*I)(nil))
	DumpMethodSet((*I1)(nil))

}
func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemTyp := v.Elem()

	n := elemTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", elemTyp)
		return
	}

	fmt.Printf("%s's method set:\n", elemTyp)

	for j := 0; j < n; j++ {
		fmt.Println("-", elemTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}
