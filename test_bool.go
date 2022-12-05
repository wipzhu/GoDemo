package main

import "fmt"

func main() {
	age := 20

	if age > 18 {
		fmt.Println("你已经成年了")
	} else {
		fmt.Println("你还未成年")
	}

	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
}
