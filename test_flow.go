package main

import "fmt"

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	x := [...]int{1, 2, 3}

	for _, v := range x {
		fmt.Println(v)
	}

}
