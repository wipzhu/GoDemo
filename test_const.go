package main

import "fmt"

func main() {
	const PI float32 = 3.14
	const PI2 = 3.1415926

	const (
		a = 100
		b = 200
	)

	const w, h = 100, 200

	// iota每调用一次，值自增1
	const (
		a1 = iota
		a2 = iota
		a3 = iota
		_
		a4 = iota
	)

	fmt.Println(PI)
	fmt.Println(PI2)
	fmt.Println(a, b, w, h)
	fmt.Println(a1, a2, a3, a4)
}
