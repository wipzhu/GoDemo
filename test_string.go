package main

import (
	"bytes"
	"fmt"
	strings "strings"
)

func main() {
	var s1 string = "hello world"

	s2 := `
	line1
	line2
	line3
	`
	fmt.Println(s1)
	fmt.Println(s2)

	//字符串拼接
	s3 := s1 + s2
	fmt.Println(s3)

	s4 := "tom"
	s5 := "20"
	s6 := fmt.Sprintf("%s,%s", s4, s5)
	fmt.Println(s6)

	name := "Jack"
	gender := "male"
	nameGender := strings.Join([]string{name, gender}, ",")
	fmt.Println(nameGender)

	var buffer bytes.Buffer
	buffer.WriteString(name)
	buffer.WriteString(",")
	buffer.WriteString(gender)
	nameGender1 := buffer.String()
	fmt.Println(nameGender1)

	//字符串切片操作
	s := "Hello World"
	a := 2
	b := 7
	fmt.Printf("%s\n", s[a:b])
	fmt.Printf("%s\n", s[a:])
	fmt.Printf("%s\n", s[0:])
	fmt.Printf("%s\n", s[:b])

	//字符串函数
	sp := strings.Split(s, " ")
	fmt.Println(sp)

	sc := strings.Contains(s, "or")
	fmt.Println(sc)

	index := strings.Index(s, "ll")
	fmt.Println(index)

	lower := strings.ToLower(s)
	fmt.Println(lower)

	upper := strings.ToUpper(s)
	fmt.Println(upper)

	prefix := strings.HasPrefix(s, "He")
	fmt.Println(prefix)

	fmt.Printf("%%")

}
