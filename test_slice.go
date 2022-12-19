package main

import (
	"fmt"
	"strconv"
)

func main() {
	test_slice()
}

func test_slice() {
	var s1 = []int{}
	var s2 = []string{}
	var s3 = make([]string, 6)
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s3)

	fmt.Println(len(s3))
	fmt.Println(cap(s3)) //容量

	lenth := len(s3)
	for i := 0; i < lenth; i++ {
		s3[i] = "wipzhu" + strconv.Itoa(i)
	}

	fmt.Printf("%v\n", s3[0:3]) // 切片取值
	fmt.Printf("%v\n", s3[2:])  // 切片取值
	fmt.Printf("%v\n", s3[:])   // 切片取值

	for i, s := range s3 {
		fmt.Printf("%v ==> %v\n", i, s) // 切片遍历
	}
}
