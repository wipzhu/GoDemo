package main

import "fmt"

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//i := 1
	//for i <= 10 {
	//	fmt.Println(i)
	//	i++
	//}

	//遍历数组
	var a = [...]int{3, 6, 9}
	for i, v := range a {
		fmt.Printf("%v => %v \n", i, v)
	}

	//遍历切片
	var b = []int{1, 2, 3}
	for _, v := range b {
		fmt.Println(v)
	}

	//遍历map
	m := make(map[string]string, 0)
	m["name"] = "Tom"
	m["age"] = "12"
	for key, value := range m {
		fmt.Printf("%v => %v \n", key, value)
	}

	//遍历字符
	s := "hello world"
	for i, v1 := range s {
		fmt.Printf("%v => %c \n", i, v1)
	}
}
