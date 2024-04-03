package main

type T struct {
	v int
}

func (t T) M1() {
	t.v = 10
}

func (t *T) M2() {
	t.v = 11
}

func main() {
	var t T
	println(t.v) // 0
	t.M1()
	println(t.v) // 0

	t.M2()
	println(t.v) // 11
}
