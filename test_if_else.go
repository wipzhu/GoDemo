package main

import (
	"fmt"
)

func main() {
	checkNum()
	checkAge()
}

func checkAge() {
	if age := 20; age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
}

func checkNum() {
	var num int
	fmt.Println("请输入一个整数:")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}
