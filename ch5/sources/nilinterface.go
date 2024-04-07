package main

// src/runtime/print.go
/*
func printeface(e eface) {
	print("(", e._type, ",", e.data, ")")
}

func printiface(i iface) {
	print("(", i.tab, ",", i.data, ")")
}
*/

func main() {
	var i interface{}
	var err error
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i = err:", i == err)
	println("")
}
