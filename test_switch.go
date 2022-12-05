package main

import "fmt"

func main() {
	//f1()
	//f2()
	//f3()

MYLABEL:

	for i := 0; i < 10; i++ {
		if i >= 5 {
			break MYLABEL
		}
		fmt.Println(i)
	}

	fmt.Println("End...")
}

func f3() {
	var score int

	fmt.Println("请输入一个数字:")
	fmt.Scan(&score)

	switch {
	case score >= 80:
		fmt.Println("优秀")
	case score >= 60 && score <= 80:
		fmt.Println("一般")
	case score < 60:
		fmt.Println("不及格")
	}
}

func f2() {
	day := 2
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	default:
		fmt.Println("非法输入")
		//没有break
	}
}
func f1() {
	grade := 'A'

	switch grade {
	case 'A':
		fmt.Println("优秀")
	case 'B':
		fmt.Println("良好")
	case 'C':
		fmt.Println("一般")
		//没有break
	}
}
