package main

func main() {
	var _name = "wipzhu"
	name := "wipzhu"
	var age int

	name, age = getNameAmdAge()

	println(_name)
	println(name)
	println(age)
}

func getNameAmdAge() (string, int) {
	return "wipzhu", 25
}
