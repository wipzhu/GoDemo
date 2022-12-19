package main

import "fmt"

func main() {
	//test_array()
	test_array2()
}

func test_array() {
	var a1 [2]int
	var a2 [7]string
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
}

func test_array2() {
	var a1 = [2]int{1, 2}
	var a2 = [2]string{"hello", "world"}
	var a3 = [2]float64{}
	var a4 = [...]string{1: "wipzhu", 3: "warmyang"}
	var a5 = [...]int{1: 2, 3: 5}

	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", a3)
	fmt.Printf("%v\n", a4)
	fmt.Printf("%v\n", a5)

	fmt.Println(len(a5))

	for i, i2 := range a4 {
		fmt.Printf("%v, ===>, %v\n", i, i2)
	}
}
