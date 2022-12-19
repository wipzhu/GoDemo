package main

import "fmt"

func main() {
	test_goto()
}

func test_goto() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i >= 2 && j >= 2 {
				goto END
			}
			fmt.Printf("%v,%v\n", i, j)
		}
	}

END:
	fmt.Println("END...")
}
